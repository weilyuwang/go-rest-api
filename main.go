package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article represents an object of an article
type Article struct {
	Title   string `json:"Title"`   // title
	Desc    string `json:"desc"`    // description
	Content string `json:"content"` // content
}

// Articles represents a list of articles of type Article
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
