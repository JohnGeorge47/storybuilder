package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JohnGeorge47/storybuilder/jsonparse"
	"github.com/JohnGeorge47/storybuilder/storyhandlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", "this")
}

func main() {
	story := jsonparse.OpenJSONFile()
	fmt.Println(story["intro"].Title)
	http.HandleFunc("/", handler)
	http.HandleFunc("/intro", storyhandlers.New(story, "intro").StoryBuilder())
	http.HandleFunc("/new-york", storyhandlers.New(story, "new-york").StoryBuilder())
	http.HandleFunc("/denver", storyhandlers.New(story, "denver").StoryBuilder())
	http.HandleFunc("/home", storyhandlers.New(story, "home").StoryBuilder())
	http.HandleFunc("/debate", storyhandlers.New(story, "debate").StoryBuilder())
	fmt.Println("Starting the server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
