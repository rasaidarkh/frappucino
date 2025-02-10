package jtoken

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	expirationJWT = time.Hour * 5
	HeaderJWT     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	/*
		{
			"alg": "HS256",
			"typ": "JWT"
		}
	*/
)

type Payload struct {
	IsAdmin   bool
	Username  string
	ExpiresAt time.Time
}

func NewPayload() *Payload {
	return &Payload{}
}

// Checks JWT format and expiration date, and returns payload
func VerifyJWT(token, secretKey string) (*Payload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	if parts[0] != HeaderJWT {
		return nil, errors.New("invalid header format")
	}

	payload := NewPayload()
	if err := json.Unmarshal([]byte(parts[1]), payload); err != nil {
		return nil, errors.New("invalid payload format")
	}

	if payload.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}

	return payload, nil
}

func SignHS256(payload []byte, secretKey string) (string, error) {
	h := hmac.New(sha256.New, []byte(secretKey))

	data, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("invalid paylaod")
	}

	h.Write([]byte(data))
	signature := h.Sum(nil)

	return base64.RawURLEncoding.EncodeToString(signature), nil
}

func GenerateAccessToken(payload *Payload, secretKey string) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", errors.New("invalid paylaod")
	}

	signature, err := SignHS256(data, secretKey)
	if err != nil {
		return "", fmt.Errorf("unable to generate access token: %v", err)
	}

	jwtToken := HeaderJWT + "." + string(data) + "." + signature

	return jwtToken, nil
}
