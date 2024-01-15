package noundo

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

const TITLE_LEN_MIN = 3
const TITLE_LEN_MAX = 50
const CONTENT_LEN_MIN = 10
const CONTENT_LEN_MAX = 5000

func validStory(story StoryContent) bool {
	return ((CONTENT_LEN_MIN <= len(story.Content)) && (len(story.Content) <= CONTENT_LEN_MAX) &&
		(TITLE_LEN_MIN <= len(story.Title)) && (len(story.Title) <= TITLE_LEN_MAX))
}

func (n *NoUndo) HandleCreateStory(w http.ResponseWriter, r *http.Request) {
	var story StoryContent

	err := json.NewDecoder(r.Body).Decode(&story)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "json_decode_fail"}, http.StatusBadRequest)
		return
	}
	jwt := GetJWTFieldsFromContext(r.Context())
	if jwt == nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "unauthorized"}, http.StatusUnauthorized)
		return
	}

	v := mux.Vars(r)

	if !validStory(story) {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "invalid_story"}, http.StatusBadRequest)
		return
	}
	history, err := n.uni.GetHistoryByName(v["history"])
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "error_during_create_story"}, http.StatusNotFound)
		return
	}
	_, err = history.CreateStory(v["age"], &User{username: strings.Split(jwt.Username, "@")[0]}, story)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "error_during_create_story"}, http.StatusNotFound)
		return
	}
	// TODO - template for a single Story
	// TODO - execute this template
	// ExecTemplHtmxSensitive(tmpl, w, r, "story", nil)
	n.HandleAge(w, r)
}
