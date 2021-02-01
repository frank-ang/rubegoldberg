package redis

// To run:
// go test -v fortune/redis
import (
	"testing"
)

func TestGetEndpoint(t *testing.T) {
	t.Log("testing getEndpoint...")
	ret := getEndpoint()
	if ret == "" {
		t.Error("Expected non-null endpoint")
	}
	t.Log("getEndpoint: " + ret)
}

func TestRandomQuote(t *testing.T) {
	t.Log("testing randomQuote...")
	ret, _ := randomQuote()
	if ret.Quote == "" {
		t.Error("Expected non-null quote")
	}
	t.Log("randomQuote: " + ret.Quote)
}
