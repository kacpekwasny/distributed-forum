package noundo

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func (n *NoUndo) HandleCreateStory(w http.ResponseWriter, r *http.Request) {
	var story StoryContent

	err := json.NewDecoder(r.Body).Decode(&story)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "json_decode_fail"}, http.StatusBadRequest)
		return
	}
	v := mux.Vars(r)
	jwt, err := JWTCheckAndParse(r)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "unauthorized"}, http.StatusUnauthorized)
		return
	}

	// TODO validate story - length, min, max

	_, err = n.Self().CreateStory(v["age"], &User{username: strings.Split(jwt.Username, "@")[0]}, story)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "error_during_create_story"}, http.StatusInternalServerError)
		return
	}
	// TODO - template for a single Story
	// TODO - execute this template
	// ExecTemplHtmxSensitive(tmpl, w, r, "story", nil)
	n.HandleAge(w, r)
}
