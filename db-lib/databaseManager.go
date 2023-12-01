package dblib

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	P "typology-processor/proto"
	M "typology-processor/structs"

	driver "github.com/arangodb/go-driver"

	"github.com/arangodb/go-driver/http"
	"github.com/go-redis/redis"
)

var Client *redis.Client
var DbClient driver.Client
var ConnectionDB driver.Database

func InitDatabases() {

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	Client = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.114:16379",
		Password: "",
		DB:       redisDB,
	})

	_, err := Client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connection to redis successful!")
	}

	// Create an HTTP connection to the database
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{os.Getenv("DATABASE_URL")},
	})
	if err != nil {
		fmt.Println(err)
	}
	// Create a client
	DbClient, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		fmt.Println(err)
	}

	ConnectionDB, err = DbClient.Database(context.Background(), "Configuration")
	if err != nil {
		fmt.Println(err)
	}
}

func AddOneGetCount(cacheKey string, data *P.FRMSMessage_Ruleresults) int64 {
	redisData, _ := json.Marshal(data)
	err := Client.SAdd(cacheKey, redisData).Err()

	if err != nil {
		fmt.Println(err)
	}

	returnValue, _ := Client.SCard(cacheKey).Result()
	return returnValue
}

func GetMembers(cacheKey string) []string {
	returnValue, _ := Client.SMembers(cacheKey).Result()
	return returnValue
}

func GetTypologyExpression(typology *P.FRMSMessage_Typologies) M.TypologyExpression {
	query := fmt.Sprintf(`FOR doc IN typologyExpression FILTER doc.id == '%s' AND doc.cfg == '%s' RETURN doc`,
		typology.Id, typology.Cfg)

	cursor, err := ConnectionDB.Query(context.Background(), query, nil)

	if err != nil {
		fmt.Print(err)
	}
	defer cursor.Close()

	var expression M.TypologyExpression
	cursor.ReadDocument(context.Background(), &expression)

	return expression
}
