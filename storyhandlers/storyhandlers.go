package storyhandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/JohnGeorge47/storybuilder/jsonparse"
)

type Info struct {
	Title   string
	Story   []string
	Options []jsonparse.Arcs
}

type Gopher struct {
	Story map[string]jsonparse.StoryChapterStruct
	Phase string
}

func New(story map[string]jsonparse.StoryChapterStruct, phase string) *Gopher {
	return &Gopher{
		Story: story,
		Phase: phase,
	}
}

//To create the story of the gopher
func (stor *Gopher) StoryBuilder() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var part string
		if r.URL.String() == "/" {
			part = "intro"
		} else {
			part = strings.TrimPrefix(stor.Phase, "/")
		}
		fmt.Println(part)
		data := Info{Title: stor.Story[part].Title, Story: stor.Story[part].Story, Options: stor.Story[part].Options}
		t, _ := template.ParseFiles("./views/intro.html")
		t.Execute(w, data)
	}
	return http.HandlerFunc(fn)
}
