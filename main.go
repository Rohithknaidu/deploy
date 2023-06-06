package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " home page end point")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8008", nil))
}
func main() {
	handleRequest()
}


