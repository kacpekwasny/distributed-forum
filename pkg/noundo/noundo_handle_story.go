package noundo

import (
	"encoding/json"
	"log/slog"
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
	return (ANSWER_LEN_MIN <= len(answer.Content)) && (len(answer.Content) <= ANSWER_LEN_MAX)
}

func (n *NoUndo) HandleCreateStoryPost(w http.ResponseWriter, r *http.Request) {
	var story StoryContent
	jwt := GetJWT(r.Context())
	if jwt == nil {
		utils.WriteJsonWithStatus(w, Unauthorized, http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&story)
	if err != nil {
		utils.WriteJsonWithStatus(w, DecodeErr, http.StatusBadRequest)
		return
	}

	if !validStory(story) {
		utils.WriteJsonWithStatus(w, InvalidValue, http.StatusBadRequest)
		return
	}

	v := mux.Vars(r)
	history, err := n.uni.GetHistoryByName(v["history"])

	if err != nil {
		utils.WriteJsonWithStatus(w, InternalError, http.StatusNotFound)
		return
	}

	_, err = history.CreateStory(jwt, v["age"], &story)
	if err != nil {
		utils.WriteJsonWithStatus(w, InternalError, http.StatusInternalServerError)
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
			AnswerToId:         storyId,
			ContentLenMin:      STORY_LEN_MIN,
			ContentLenMax:      STORY_LEN_MAX,
			WriteAnswerPostURL: WriteAnswerURL(historyName, storyId),
			HideAfterSend:      false,
		},
	})
}

func (n *NoUndo) HandleCreateAnswerGet(w http.ResponseWriter, r *http.Request) {
	jwt := GetJWT(r.Context())
	if jwt == nil {
		// TODO
		http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		return
	}

	currURL := r.Header.Get("hx-current-url")
	URL, err := url.Parse(currURL)
	if err != nil {
		utils.WriteJsonWithStatus(w, InvalidHeaders, http.StatusBadRequest)
		return
	}
	parts := strings.Split(URL.Path, "/")
	if len(parts) < 5 {
		utils.WriteJsonWithStatus(w, InvalidURL, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	postableId := vars["postable-id"]

	tmpl.ExecuteTemplate(w, "answer_write", CompAnswerWrite{
		AnswerToId:         mux.Vars(r)["postable-id"],
		ContentLenMin:      ANSWER_LEN_MIN,
		ContentLenMax:      ANSWER_LEN_MAX,
		WriteAnswerPostURL: WriteAnswerURL(parts[2], postableId),
		HideAfterSend:      true,
	})

}

func (n *NoUndo) HandleCreateAnswerPost(w http.ResponseWriter, r *http.Request) {
	jwt := GetJWT(r.Context())
	if jwt == nil {
		http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		return
	}
	// TODO get the POST body of answer

	var answerContent AnswerContent
	err := json.NewDecoder(r.Body).Decode(&answerContent)
	if err != nil {
		utils.WriteJsonWithStatus(w, InvalidValue, http.StatusBadRequest)
		return
	}
	if !validAnswer(answerContent) {
		slog.Debug("invalid_answer", "answerContent", answerContent)
		utils.WriteJsonWithStatus(w, InvalidValue, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	historyName := vars["history"]
	parentId := vars["postable-id"]
	history, err := n.uni.GetHistoryByName(historyName)
	if err != nil {
		utils.WriteJsonWithStatus(w, NotFound, http.StatusNotFound)
		return
	}

	answer, err := history.CreateAnswer(jwt, parentId, answerContent.Content)
	if err != nil {
		utils.WriteJsonWithStatus(w, InternalError, http.StatusInternalServerError)
		return
	}
	utils.ExecTemplLogErr(tmpl, w, "answer_tree_node", answer)
}
