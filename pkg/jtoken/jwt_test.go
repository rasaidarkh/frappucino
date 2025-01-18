package jtoken_test

import (
	"context"
	"frappuccino/pkg/jtoken"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestCreateJWT(t *testing.T) {
	server, _ := miniredis.Run()

	token, err := jtoken.GenerateAccessToken(
		context.Background(),
		redis.NewClient(&redis.Options{Addr: server.Addr()}),
		map[string]interface{}{},
	)

	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}
