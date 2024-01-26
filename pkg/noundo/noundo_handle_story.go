package noundo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

const TITLE_LEN_MIN = 3
const TITLE_LEN_MAX = 50

const STORY_LEN_MIN = 10
const STORY_LEN_MAX = 5000

const ANSWER_LEN_MIN = 3
const ANSWER_LEN_MAX = STORY_LEN_MAX

func validStory(story StoryContent) bool {
	return ((STORY_LEN_MIN <= len(story.Content)) && (len(story.Content) <= STORY_LEN_MAX) &&
		(TITLE_LEN_MIN <= len(story.Title)) && (len(story.Title) <= TITLE_LEN_MAX))
}

func validAnswer(answer AnswerContent) bool {
	return (STORY_LEN_MIN <= len(answer.Content)) && (len(answer.Content) <= STORY_LEN_MAX)
}

func (n *NoUndo) HandleCreateStoryPost(w http.ResponseWriter, r *http.Request) {
	var story StoryContent
	jwt := GetJWT(r.Context())
	if jwt == nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "unauthorized"}, http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&story)
	if err != nil {
		utils.WriteJsonWithStatus(w, utils.Ms{"info": "json_decode_fail"}, http.StatusBadRequest)
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

func (n *NoUndo) HandleStoryGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	historyName := vars["history"]
	storyId := vars["story-id"]

	histIface, err := n.uni.GetHistoryByName(historyName)
	if err != nil {
		n.Handle404(w, r)
		return
	}

	story, err := histIface.GetStory(storyId)
	if err != nil {
		n.Handle404(w, r)
		return
	}

	ExecTemplHtmxSensitive(tmpl, w, r, "story_page", r.URL.Path, &PageStoryValues{
		PageBaseValues:      CreatePageBaseValues(story.Title, n.Self(), histIface, r),
		CompAgeHeaderValues: CreateAgeHeader(historyName, story.AgeName),
		CompStoryValues: CompStoryValues{
			Story:        story,
			ClampContent: false,
			StoryURL:     "",
		},
		CompAnswerWrite: CompAnswerWrite{
			AnswerToId:    storyId,
			ContentLenMin: STORY_LEN_MIN,
			ContentLenMax: STORY_LEN_MAX,
		},
	})
}

func (n *NoUndo) HandleCreateAnswerGet(w http.ResponseWriter, r *http.Request) {
	jwt := GetJWT(r.Context())
	if jwt == nil {
		http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
		return
	}

	currURL := r.Header.Get("hx-current-url")
	URL, err := url.Parse(currURL)
	if err != nil {
		utils.WriteJsonWithStatus(w, "headers_not_matching_requirements", http.StatusBadRequest)
		return
	}
	parts := strings.Split(URL.Path, "/")
	if len(parts) < 5 {
		utils.WriteJsonWithStatus(w, "expected_url_to_contain_more_information", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	postableId := vars["postable-id"]

	tmpl.ExecuteTemplate(w, "answer_write", CompAnswerWrite{
		AnswerToId:         mux.Vars(r)["postable-id"],
		ContentLenMin:      ANSWER_LEN_MIN,
		ContentLenMax:      ANSWER_LEN_MAX,
		WriteAnswerPostURL: utils.LeftLogRight(url.JoinPath("/write-answer", parts[2], postableId)),
	})

}

func (n *NoUndo) HandleCreateAnswerPost(w http.ResponseWriter, r *http.Request) {
	jwt := GetJWT(r.Context())
	if jwt == nil {
		http.Redirect(w, r, "/sign-in", http.StatusTemporaryRedirect)
		return
	}
	// TODO get the POST body of answer

	var answerContent AnswerContent
	err := json.NewDecoder(r.Body).Decode(&answerContent)
	if err != nil {
		utils.WriteJsonWithStatus(w, "parsing_body_failed", http.StatusBadRequest)
		return
	}
	if validAnswer(answerContent) {
		utils.WriteJsonWithStatus(w, "invalid_answer", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	historyName := vars["history"]
	parentId := vars["postable-id"]
	history, err := n.uni.GetHistoryByName(historyName)
	if err != nil {
		utils.WriteJsonWithStatus(w, "history_not_found", http.StatusNotFound)
		return
	}

	history.GetUser(jwt.Username)
	answer, err := history.CreateAnswer(&User{username: jwt.Username, parentServerName: jwt.ParentServer}, parentId, answerContent.Content)
	if err != nil {
		utils.WriteJsonWithStatus(w, "error_creating_answer", http.StatusInternalServerError)
		return
	}
	utils.ExecTemplLogErr(tmpl, w, "answer", answer)
}
