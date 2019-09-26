package storyhandlers

import (
	"net/http"

	"github.com/JohnGeorge47/storybuilder/jsonparse"
)

func Introhandler(story map[string]jsonparse.StoryChapterStruct) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

	}
	return http.HandlerFunc(fn)
}
