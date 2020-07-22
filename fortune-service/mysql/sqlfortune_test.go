package mysql

// To run:
// go test -v fortune/mysql
import (
	"fmt"
	"testing"
)

func TestMySqlGetEndpoint(t *testing.T) {
	t.Log("testing getMySqlEndpoint...")
	ret := getMySqlEndpoint()
	if ret == "" {
		t.Error("Expected non-null endpoint")
	}
	t.Log("sql endpoint: " + ret)
}

func TestMySqlGetSecrets(t *testing.T) {
	t.Log("testing secrets manager get creds...")
	username, password, _ := getMySqlCredentials()
	if username == "" {
		t.Error("Expected non-null username")
	}
	if password == "" {
		t.Error("Expected non-null password")
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
