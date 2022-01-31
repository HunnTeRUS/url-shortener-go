package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func addLink(w http.ResponseWriter, r *http.Request) {
	log.Println("Add Link")
	key, ok := r.URL.Query()["link"]
	if ok {
		if !validLink(key[0]) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Could not create shortlink need absolute path link. Ex: /addLink?link=https://github.com/")
			return
		}
		log.Println(key)
		if _, ok := linkList[key[0]]; !ok {
			genString := randStringBytes(10)
			linkList[genString] = key[0]
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusAccepted)

			linkString := fmt.Sprintf("<a href=\"http://localhost:9000/short/%s\">http://localhost:9000/short/%s</a>", genString, genString)
			fmt.Fprintf(w, "Added shortlink\n")
			fmt.Fprintf(w, linkString)
			return
		}
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Already have this link")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Failed to add link")
	return
}

// validLink - check that the link we're creating a shortlink for is a absolute URL path
func validLink(link string) bool {
	r, err := regexp.Compile("^(http|https)://")
	if err != nil {
		return false
	}
	link = strings.TrimSpace(link)
	log.Printf("Checking for valid link: %s", link)
	// Check if string matches the regex
	if r.MatchString(link) {
		return true
	}
	return false
}
