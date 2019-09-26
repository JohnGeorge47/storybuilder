package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnGeorge47/storybuilder/jsonparse"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", "this")
}

func main() {
	story := jsonparse.OpenJSONFile()
	fmt.Println(story["intro"].Title)
	http.HandleFunc("/", handler)
	// http.HandleFunc("/edit", editHandler)
	fmt.Println("Starting the server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
