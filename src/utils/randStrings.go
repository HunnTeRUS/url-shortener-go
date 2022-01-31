package utils

import (
	"fmt"
	"math/rand"

	"github.com/HunnTeRUS/url-shortener-go/src/configuration/env"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	BASE_URL = "BASE_URL"
)

// randStringBytes - Create random short link
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateURL() string {
	genString := randStringBytes(11)

	base_url := env.GetEnvVar(BASE_URL, "http://localhost:9000")
	linkString := fmt.Sprintf("%s/short/%s", base_url, genString)

	return linkString
}
