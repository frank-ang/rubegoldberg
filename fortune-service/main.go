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
)

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
	http.Handle("/", r)

	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
	fmt.Println("Exiting.")
}

func Greeting(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, Fortune!")
}
