package mysql

// To run:
// go test -v fortune/mysql
import (
	"fmt"
	"testing"
)

func TestMySqlGetSecrets(t *testing.T) {
	t.Log("testing secrets manager get creds...")
	username, password, db_host, _ := getMySqlConnectionInfo()
	if username == "" || password == "" || db_host == "" {
		t.Error("Encountered null username, password, or db_host")
	}
}

func TestMySqlFortune(t *testing.T) {
	t.Log("testing init ...")
	Init()
	count, _ := queryQuoteCount()
	if count == 1 {
		t.Error(fmt.Sprintf("count is incorrect: %d", count))
	}
	t.Logf("count: %d", count)
	quote := randomQuote()
	t.Logf("quote: %+v", quote)
}
