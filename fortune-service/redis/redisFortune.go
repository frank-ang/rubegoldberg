package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	REDIS_HOST = os.Getenv("REDIS_HOST")
)

type Quote struct {
	Id                   int
	Quote, Author, Genre string
}

func GetFortune(w http.ResponseWriter, req *http.Request) {
	log.Print("Getting quote.")
	quote := randomQuote()
	jsonQuote, err := json.MarshalIndent(&quote, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonQuote))
}

func randomQuote() Quote {
	var q Quote
	redisClient := getRedisClient()
	var ctx = context.Background()
	id := 1 // hardcode test.
	key := fmt.Sprint("quote:", id)
	val, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	q.Id = id
	q.Author = val["author"]
	q.Genre = val["genre"]
	q.Quote = val["quote"]
	return q
}

func getRedisClient() *redis.Client {
	log.Printf("Getting Client for Redis host: %s", getEndpoint())
	client := redis.NewClient(&redis.Options{
		Addr:     getEndpoint(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func getEndpoint() string {
	return REDIS_HOST + ":6379"
}
