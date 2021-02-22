package main

import (
	"flag"
	"fmt"
	"fortune/elasticsearch"
	"fortune/mysql"
	"fortune/redis"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
)

func init() {
	fmt.Println("fortune main init()")
	xray.Configure(xray.Config{
		ServiceVersion: "0.0.1",
	})
}

// Fortune. Returns a famous quote..
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	mysqlPtr := flag.Bool("mysql", false, "mysql handler")
	esPtr := flag.Bool("es", false, "elasticsearch handler")
	redisPtr := flag.Bool("redis", false, "redis handler")

	flag.Parse()
	log.Printf("Flags: mysql=%t, redis=%t, elasticsearch=%t", *mysqlPtr, *redisPtr, *esPtr)

	http.Handle("/",
		xray.Handler(xray.NewFixedSegmentNamer("fortune.root"),
			http.HandlerFunc(Greeting)))
	http.Handle("/health",
		xray.Handler(xray.NewFixedSegmentNamer("fortune.health"),
			http.HandlerFunc(Greeting)))
	if *mysqlPtr {
		http.Handle("/fortune/sql",
			xray.Handler(xray.NewFixedSegmentNamer("fortune.sql"),
				http.HandlerFunc(mysql.GetFortune)))
	}
	if *esPtr {
		http.Handle("/fortune/es",
			xray.Handler(xray.NewFixedSegmentNamer("fortune.es"),
				http.HandlerFunc(elasticsearch.GetFortune)))
	}
	if *redisPtr {
		http.Handle("/fortune/redis",
			xray.Handler(xray.NewFixedSegmentNamer("fortune.redis"),
				http.HandlerFunc(redis.GetFortune)))
	}
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("Exiting.")
}

func Greeting(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Fprintln(w, "{ \"message\": \"Hello, Fortune!\" }")
}
