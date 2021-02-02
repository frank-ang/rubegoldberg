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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

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

	// TODO remove
	//http.Handle("/", xray.Handler(xray.NewFixedSegmentNamer("fortune"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello fortune from xray!"))
	//})))
	// or...
	// router.Handle("/", xray.Handler(xray.NewFixedSegmentNamer("newsegment"), HelloServer()))

	r := mux.NewRouter()
	r.HandleFunc("/", Greeting)
	r.HandleFunc("/health", Greeting)
	if *mysqlPtr {
		r.HandleFunc("/fortune/sql", mysql.GetFortune)
	}
	if *esPtr {
		r.HandleFunc("/fortune/es", elasticsearch.GetFortune)
	}
	if *redisPtr {
		r.HandleFunc("/fortune/redis", redis.GetFortune)
	}
	// TODO: xrayHandler, doesn't appear to be handling anything?
	// xrayHandler := xray.Handler(xray.NewFixedSegmentNamer("fortune"), r)
	http.Handle("/", r)

	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
	fmt.Println("Exiting.")
}

func Greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, Fortune!")
}
