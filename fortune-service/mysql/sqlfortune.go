package mysql

/*
Golang noob MVP implementation of fortune lookup against MySQL DB.

TODO: Retry connection errors in Secrets lookups and connection pool.
*/
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/avast/retry-go"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

var (
	initialized    bool = false
	DB_SECRET_NAME      = os.Getenv("DB_SECRET_NAME")
	DB_HOST             = os.Getenv("DB_HOST")
	username       string
	password       string
	db             *sql.DB
)

type Quote struct {
	Id                   int
	Quote, Author, Genre string
}

func GetFortune(w http.ResponseWriter, req *http.Request) {
	log.Print("Getting quote.")

	if !initialized {
		retry.Do(
			func() error {
				return Init()
			},
		)
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

func Init() error {
	username, password, DB_HOST, err := getMySqlConnectionInfo()
	if err != nil {
		return err
	}
	err2 := initConnectionPool(username, password, DB_HOST)
	if err2 != nil {
		return err2
	}
	log.Printf("Verifying DB connection to: %s", DB_HOST)
	// err := db.Ping() // getting timeouts on aurora somehow..?
	count, err3 := queryQuoteCount()
	if err3 != nil {
		fmt.Println(err3.Error())
		return err3
	}
	log.Printf("initialization complete. count of records: %d", count)
	initialized = true
	return nil
}

func initConnectionPool(username string, password string, endpoint string) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/demo", username, password, endpoint)
	log.Printf("initializing connection pool...")
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(5)
	return nil
}

func getMySqlConnectionInfo() (string, string, string, error) {
	if username != "" && password != "" && DB_HOST != "" {
		return username, password, DB_HOST, nil
	}

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
		return "", "", "", err
	}
	valueMap := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(*output.SecretString), &valueMap)

	if err2 != nil {
		panic(err2)
	}
	secretValue := valueMap["password"].(string)
	username := valueMap["username"].(string)
	dbHost := valueMap["host"].(string)
	if DB_HOST != "" {
		dbHost = DB_HOST
	}
	return username, secretValue, dbHost, nil
}
