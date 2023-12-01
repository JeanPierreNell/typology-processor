package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	db "typology-processor/db-lib"
	P "typology-processor/proto"

	M "typology-processor/structs"

	"golang.org/x/exp/slices"
)

func HandleTransaction(message *P.FRMSMessage) {
	networkmap := message.NetworkMap
	ruleResult := message.RuleResult
	transaction := message.Transaction

	for _, channel := range networkmap.Messages[0].Channels {
		for _, typology := range channel.Typologies {
			executeRequest(transaction, typology, ruleResult, networkmap)
		}
	}
}

func executeRequest(transaction *P.FRMSMessage_Transaction, typology *P.FRMSMessage_Typologies, ruleResult *P.FRMSMessage_Ruleresults, networkMap *P.FRMSMessage_Networkmap) {
	transactionID := transaction.FIToFIPmtSts.GrpHdr.MsgId
	cacheKey := fmt.Sprintf("TP_%s_%s_%s", transactionID, typology.Id, typology.Cfg)
	jruleResultsCount := db.AddOneGetCount(cacheKey, ruleResult)

	if int(jruleResultsCount) < len(typology.Rules) {
		return
	}

	jruleResults := db.GetMembers(cacheKey)
	ruleResults := make([]*P.FRMSMessage_Ruleresults, 0)

	for _, ruleResult := range jruleResults {
		currentRule := P.FRMSMessage_Ruleresults{}
		json.Unmarshal([]byte(ruleResult), &currentRule)
		ruleResults = append(ruleResults, &currentRule)
	}

	typologyExpression := db.GetTypologyExpression(typology)
	typologyResultValue := evaluateTypologyExpression(typologyExpression.Rules, ruleResults, typologyExpression)
	fmt.Print(typologyResultValue)

	typologyResult := P.FRMSMessage_Typologyresult{}
	typologyResult.Result = typologyResultValue
	typologyResult.RuleResults = ruleResults
	typologyResult.Cfg = typology.Cfg
	typologyResult.Id = typology.Id

	//Send Response to NATS
	HandleResponse(&typologyResult)
}

func evaluateTypologyExpression(ruleValues []M.RuleConfig, ruleResults []*P.FRMSMessage_Ruleresults, typologyExpression M.TypologyExpression) uint32 {
	var toReturn uint32 = 0
	for _, rule := range ruleResults {
		ruleResult := ruleResults[slices.IndexFunc(ruleResults, func(r *P.FRMSMessage_Ruleresults) bool {
			return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
		})]

		var ruleVal uint64 = 0

		if ruleResult.Result {
			ruleVal, _ = strconv.ParseUint(ruleValues[slices.IndexFunc(ruleValues, func(v M.RuleConfig) bool {
				return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
			})].True, 10, 32) //Add Stuff / Rework
		} else {
			ruleVal, _ = strconv.ParseUint(ruleValues[slices.IndexFunc(ruleValues, func(v M.RuleConfig) bool {
				return rule.Id == typologyExpression.Expression.Terms[0].Id && rule.Cfg == typologyExpression.Expression.Terms[0].Cfg
			})].False, 10, 32) //Find way to direct parse to uint32
		}

		switch typologyExpression.Expression.Operator {
		case "+":
			toReturn += uint32(ruleVal)
		case "-":
			toReturn -= uint32(ruleVal)
		case "*":
			toReturn *= uint32(ruleVal)
		case "/":
			toReturn /= uint32(ruleVal)
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
