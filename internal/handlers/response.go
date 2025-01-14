package handlers

import "context"

type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

func JSON(ctx context.Context, status int, data interface{}) {}
