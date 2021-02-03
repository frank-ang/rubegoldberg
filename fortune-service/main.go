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

// Fortune. Returns a famous quote.
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
		xray.Handler(xray.NewFixedSegmentNamer("fortune"),
			http.HandlerFunc(Greeting)))
	http.Handle("/health",
		xray.Handler(xray.NewFixedSegmentNamer("fortune"),
			http.HandlerFunc(Greeting)))
	if *mysqlPtr {
		http.Handle("/fortune/sql",
			xray.Handler(xray.NewFixedSegmentNamer("fortune"),
				http.HandlerFunc(mysql.GetFortune)))
	}
	if *esPtr {
		http.Handle("/fortune/es",
			xray.Handler(xray.NewDynamicSegmentNamer("fortune.es", "*.amazonaws.com"),
				http.HandlerFunc(elasticsearch.GetFortune)))
	}
	if *redisPtr {
		http.Handle("/fortune/redis",
			xray.Handler(xray.NewDynamicSegmentNamer("fortune.redis", "*.amazonaws.com"),
				http.HandlerFunc(redis.GetFortune)))
	}
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	fmt.Println("Exiting.")
}

func Greeting(w http.ResponseWriter, req *http.Request) {
	_, seg := xray.BeginSegment(req.Context(), "fortune.greeting")
	fmt.Fprintln(w, "Hello, Fortune!")
	seg.Close(nil)
}
