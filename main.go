package main

import (
  "fmt"
  "log"
  "net/http" // we can also use the gorilla/mux also for the request instead of net/http
  "encoding/json"
  "github.com/gorilla/mux" // we can use the base url for the package because it's an open source packages.
)

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Homepage Endpoint Hit")
  fmt.Println("Homepage Endpoint Hit")
}

func returnAllArticle(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Endpoint Hit: returnAllArticle")
  json.NewEncoder(w).Encode(Articles) // encode the struct that we will be served as a json.
}

// using the net/http
func handleRequests() {
  http.HandleFunc("/api", homePage) // the route of the page {{url:8081}/api}
  http.HandleFunc("/api/article", returnAllArticle) // the route of the page {{url:8081}/api/article}
  log.Fatal(http.ListenAndServe(":8081", nil))
}


type Article struct {
  Title string `json: "title"`
  Desc string `json: "desc"`
  Content string `json: "content"`
}

// declare a global array variables of struct Article
var Articles []Article 


func main() {
  // populate the array
  Articles = []Article {
    Article{Title: "Hello", Desc: "This is the Description for The first element of the array", Content: "This is the First Article Content"},
    Article{Title: "Hello", Desc: "This is the Description for The second element of the array", Content: "This is the Second Article Content"},
  }
  handleRequests() // call the handle request that'll call 
                   // the Homepage Endpoint Hit of route /
}
