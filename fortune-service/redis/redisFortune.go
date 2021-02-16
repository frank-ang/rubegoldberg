package redis

/*
Basic fortune lookup against Redis.
TODO: retries
*/
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

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
	quote, quoteErr := randomQuote()
	if quoteErr != nil {
		fmt.Println(quoteErr)
		http.Error(w, quoteErr.Error(), 500)
		return
	}
	jsonQuote, err := json.MarshalIndent(&quote, "", "    ")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	fmt.Fprintln(w, string(jsonQuote))
}

func randomQuote() (Quote, error) {
	var q Quote
	redisClient := getRedisClient()
	var ctx = context.Background()
	size64, sizeErr := redisClient.DBSize(ctx).Result()
	if sizeErr != nil {
		return q, sizeErr
	}
	size := int(size64)
	rand.Seed(time.Now().UnixNano())
	id := int(rand.Intn(size))
	log.Print("DBSize: ", size, ", random ID:", id)
	key := fmt.Sprint("quote:", id)
	val, err := redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return q, err
	}
	q.Id = id
	q.Author = val["author"]
	q.Genre = val["genre"]
	q.Quote = val["quote"]
	return q, nil
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
