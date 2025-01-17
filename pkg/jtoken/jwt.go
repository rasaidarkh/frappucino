package jtoken

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"frappuccino/pkg/config"
	"strings"

	"github.com/go-redis/redis/v8"
)

const (
	HeaderJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	/*
		{
			"alg": "HS256",
			"typ": "JWT"
		}
	*/
)

func VerifyJWT(ctx context.Context, rdb *redis.Client, token, secretKey string) (bool, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false, errors.New("invalid token format")
	}

	err := rdb.Get(ctx, token).Err()
	if err == redis.Nil {
		return false, fmt.Errorf("token %s does not exist", token)
	} else if err != nil {
		return false, fmt.Errorf("unable to fetch token from redis")
	}

	return true, nil
}

func SignHS256(payload map[string]interface{}, secretKey string) (string, error) {
	h := hmac.New(sha256.New, []byte(secretKey))

	data, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("invalid paylaod")
	}

	h.Write([]byte(data))
	signature := h.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(signature), nil
}

func GenerateAccessToken(ctx context.Context, rdb *redis.Client, payload map[string]interface{}) (string, error) {
	cfg := config.GetConfing()

	payloadToken, err := SignHS256(payload, cfg.JWTSecret)
	if err != nil {
		return "", fmt.Errorf("unable to generate access token: %v", err)
	}

	jwtToken := HeaderJWT + "." + payloadToken + "." + cfg.JWTSecret

	err = rdb.HSet(ctx, jwtToken, payload).Err()
	if err != nil {
		return "", fmt.Errorf("unable to set key in redis: %v", err)
	}

	return jwtToken, nil
}
