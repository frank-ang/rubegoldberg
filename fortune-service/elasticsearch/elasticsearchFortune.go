package elasticsearch

/*
Basic fortune lookup against ElasticSearch.
TODO: retry connection errors.
*/

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

var (
	ES_HOST = os.Getenv("ES_HOST")
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
	log.Printf("looking up quote...")
	elasticClient := getElasticClient()
	log.Printf("Elasticsearch version: %s", elasticsearch.Version)
	var query = `{
		"size": 1,
		"query": {
			"function_score": {
				"query": { "match_all": {} },
				"functions": [ 
					{
						"random_score": {} 
					}
				]
			}
		}
	}
	`

	// Instantiate a mapping interface for API response
	var mapResp map[string]interface{}
	var buf bytes.Buffer
	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	// Query is a valid JSON object
	res, err := elasticClient.Search(
		elasticClient.Search.WithContext(context.Background()),
		elasticClient.Search.WithIndex("quotes"),
		elasticClient.Search.WithBody(read),
		elasticClient.Search.WithTrackTotalHits(true),
		elasticClient.Search.WithPretty(),
	)

	// Check for any errors returned by API call to Elasticsearch
	if err != nil {
		log.Fatal("Elasticsearch Search() API ERROR:", err)
	}
	// If no errors are returned, parse esapi.Response object
	// Close the result body when the function call is complete
	defer res.Body.Close()

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(res.Body).Decode(&mapResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// If no error, then convert response to a map[string]interface
	// Iterate the document "hits" returned by API call
	for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {
		// Parse the attributes/fields of the document
		doc := hit.(map[string]interface{})
		// The "_source" data is another map interface nested inside of doc
		source := doc["_source"].(map[string]interface{})
		fmt.Println("doc _source:", reflect.TypeOf(source))
		// Get the doc values
		//fmt.Println("_source:", source)
		q.Id, _ = strconv.Atoi(doc["_id"].(string))
		q.Author = source["author"].(string)
		q.Genre = source["genre"].(string)
		q.Quote = source["quote"].(string)
		break
	}

	return q
}

func getElasticClient() *elasticsearch.Client {
	// init Elastic client
	log.Printf("Getting Client for Elasticsearch host: %s", getEndpoint())
	cfg := elasticsearch.Config{
		Addresses: []string{
			getEndpoint(),
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	elasticClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("Elasticsearch connection error:", err)
	}
	return elasticClient
}

func getEndpoint() string {
	return ES_HOST
}
