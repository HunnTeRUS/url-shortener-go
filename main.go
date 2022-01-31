package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	linkList map[string]string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	linkList = map[string]string{}

	http.HandleFunc("/addLink", addLink)
	http.HandleFunc("/short/", getLink)
	http.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":9000", nil))
}
