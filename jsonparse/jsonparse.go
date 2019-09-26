package jsonparse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//StoryChapterStruct is base for each story
type StoryChapterStruct struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Arcs   `json:"options"`
}

//Arcs define each option for user to choose
type Arcs struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

//ParseIt parses the json of the story
func ParseIt(storymap map[string]StoryChapterStruct) {

}

//OpenJSONFile the json file from filesystem
func OpenJSONFile() map[string]StoryChapterStruct {
	filepath := "/home/exotel/go/src/github.com/JohnGeorge47/storybuilder/jsonparse/gopher.json"
	fileopen, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	var jsonMap map[string]StoryChapterStruct
	file, err := ioutil.ReadAll(fileopen)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal([]byte(file), &jsonMap)
	if err != nil {
		fmt.Println(err)
	}
	return jsonMap
}
