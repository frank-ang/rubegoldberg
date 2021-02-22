package redis

// To run:
// go test -v fortune/redis
import (
	"context"
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
	var ctx = context.Background()
	ret, _ := randomQuote(ctx)
	if ret.Quote == "" {
		t.Error("Expected non-null quote")
	}
	t.Log("randomQuote: " + ret.Quote)
}
