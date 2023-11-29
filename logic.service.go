package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	db "typology-processor/db-lib"
	M "typology-processor/structs"

	"golang.org/x/exp/slices"
)

func HandleTransaction(message string) {
	parsedMessage := M.TransactionRequest{}
	err := json.Unmarshal([]byte(message), &parsedMessage)

	if err != nil {
		fmt.Println(err)
		return
	}

	networkmap := parsedMessage.NetworkMap
	ruleResult := parsedMessage.RuleResult
	transaction := parsedMessage.Transaction

	for _, channel := range networkmap.Messages[0].Channels {
		for _, typology := range channel.Typologies {
			executeRequest(transaction, typology, ruleResult, networkmap)
		}
	}
}

func executeRequest(transaction M.Pain002, typology M.Typology, ruleResult M.RuleResult, networkMap M.NetworkMap) {
	transactionID := transaction.FIToFIPmtSts.GrpHdr.MsgId
	cacheKey := fmt.Sprintf("TP_%s_%s_%s", transactionID, typology.Id, typology.Cfg)
	jruleResultsCount := db.AddOneGetCount(cacheKey, ruleResult)

	if int(jruleResultsCount) < len(typology.Rules) {
		return
	}

	jruleResults := db.GetMembers(cacheKey)
	ruleResults := make([]M.RuleResult, 0)

	for _, ruleResult := range jruleResults {
		currentRule := M.RuleResult{}
		json.Unmarshal([]byte(ruleResult), &currentRule)
		ruleResults = append(ruleResults, currentRule)
	}

	typologyExpression := db.GetTypologyExpression(typology)
	typologyResultValue := evaluateTypologyExpression(typologyExpression.Rules, ruleResults, typologyExpression)
	fmt.Print(typologyResultValue)

	typologyResult := M.TypologyResult{}
	typologyResult.Result = typologyResultValue
	typologyResult.RuleResult = ruleResults
	typologyResult.Cfg = typology.Cfg
	typologyResult.Id = typology.Id

	//Send Response to NATS
	HandleResponse(typologyResult)
}

func evaluateTypologyExpression(ruleValues []M.RuleConfig, ruleResults []M.RuleResult, typologyExpression M.TypologyExpression) float64 {
	toReturn := 0.0
	for _, rule := range ruleResults {
		ruleResult := ruleResults[slices.IndexFunc(ruleResults, func(r M.RuleResult) bool {
			return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
		})]
		ruleVal := 0.0

		if ruleResult.Result {
			ruleVal, _ = strconv.ParseFloat(ruleValues[slices.IndexFunc(ruleValues, func(v M.RuleConfig) bool {
				return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
			})].True, 1) //Add Stuff / Rework
		} else {
			ruleVal, _ = strconv.ParseFloat(ruleValues[slices.IndexFunc(ruleValues, func(v M.RuleConfig) bool {
				return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
			})].False, 1)
		}

		switch typologyExpression.Expression.Operator {
		case "+":
			toReturn += ruleVal
		case "-":
			toReturn -= ruleVal
		case "*":
			toReturn *= ruleVal
		case "/":
			toReturn /= ruleVal
		}
	}
	// if typologyExpression.Expression.Operator != "" {
	// 	evalResult := evaluateTypologyExpression(ruleValues, ruleResults, typologyExpression)
	// 	switch typologyExpression.Expression.Operator {
	// 	case "+":
	// 		toReturn += evalResult
	// 	case "-":
	// 		toReturn -= evalResult
	// 	case "*":
	// 		toReturn *= evalResult
	// 	case "/":
	// 		toReturn /= evalResult
	// 	}
	// } Some sort of Recursion, don't see sample in sample data as this moment

	return toReturn
}
