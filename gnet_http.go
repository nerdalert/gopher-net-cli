package main

import (
	"net/http"
	"time"
)

func NewClient() http.Client {
	timeout := time.Duration(5 * time.Second)
	return http.Client{
		Timeout: timeout,
	}
}
