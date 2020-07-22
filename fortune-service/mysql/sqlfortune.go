package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

var (
	initialized    bool = false
	DB_SECRET_NAME      = os.Getenv("DB_SECRET_NAME")
	DB_HOST             = os.Getenv("DB_HOST")
	db             *sql.DB
)

type Quote struct {
	Id                   int
	Quote, Author, Genre string
}

func GetFortune(w http.ResponseWriter, req *http.Request) {
	log.Print("Getting quote.")
	if !initialized {
		Init()
		initialized = true
	}
	quote := randomQuote()
	jsonQuote, err := json.MarshalIndent(&quote, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintln(w, string(jsonQuote))
}

func randomQuote() Quote {
	query := "SELECT * FROM quotes WHERE id IN (SELECT id FROM (SELECT id FROM quotes ORDER BY RAND() LIMIT 1) t)"
	var q Quote
	err := db.QueryRow(query).Scan(&q.Id, &q.Quote, &q.Author, &q.Genre)
	if err != nil {
		log.Panicf("query error: %v\n", err)
	}
	return q
}

func queryQuoteCount() (int, error) {
	count := 0
	err := db.QueryRow("SELECT count(1) FROM quotes").Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return count, nil
}

func Init() {
	username, password, _ := getMySqlCredentials()
	DB_HOST := getMySqlEndpoint()
	initConnectionPool(username, password, DB_HOST)
	log.Printf("Verifying DB connection to: %s", DB_HOST)
	// err := db.Ping() // getting timeouts on aurora somehow..?
	count, err := queryQuoteCount()
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("initialization complete. count of records: %d", count)
}

func initConnectionPool(username string, password string, endpoint string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/demo", username, password, endpoint)
	log.Printf("initializing connection pool...")
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxIdleConns(5)
}

func getMySqlCredentials() (string, string, error) {
	log.Printf("getting sql creds with secret name: %s", DB_SECRET_NAME)
	sess := session.Must(session.NewSession())
	secretsmanagerSvc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &DB_SECRET_NAME,
	}
	log.Printf("GetSecretValueInput: %+v", input)

	log.Print("looking up secrets manager...")
	output, err := secretsmanagerSvc.GetSecretValue(input)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}
	valueMap := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(*output.SecretString), &valueMap)

	if err2 != nil {
		panic(err2)
	}
	secretValue := valueMap["password"].(string)
	username := valueMap["username"].(string)

	return username, secretValue, nil
}

// Get the MySQL endpoint from: Secrets Manager, or DB_HOST env variable
func getMySqlEndpoint() string {
	log.Print("looking up database endpoint...")
	if DB_HOST != "" {
		return DB_HOST
	}
	log.Printf("getting endpoint from secret: %s", DB_SECRET_NAME)
	sess := session.Must(session.NewSession())
	secretsmanagerSvc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &DB_SECRET_NAME,
	}
	output, err := secretsmanagerSvc.GetSecretValue(input)
	if err != nil {
		fmt.Println(err.Error())
		panic(err) // TODO should retry Secrets Manager lookup.
	}
	valueMap := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(*output.SecretString), &valueMap)
	if err2 != nil {
		panic(err2)
	}
	endpoint := valueMap["host"].(string)
	return endpoint
}
