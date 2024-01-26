package noundo

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func (n *NoUndo) HandleAge(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	historyName := params["history"]
	ageName := params["age"]
	history, err := n.uni.GetHistoryByName(historyName)
	if err != nil {
		n.Handle404(w, r)
		return
	}

	_, err = history.GetAge(ageName)
	if err != nil {
		// TODO (create Age page)
		n.Handle404(w, r)
		return
	}

	stories, err := history.GetStories(
		[]string{ageName},
		int(utils.GetQueryParamInt(r, "start", 0)),
		int(utils.GetQueryParamInt(r, "end", 50)),
		nil, nil,
	)

	if err != nil {
		// TODO, logging, user info, maybe CreateAge option?
		n.HandleHome(w, r)
		return
	}

	storiesForTmpl := make([]CompStoryValues, len(stories))
	for i, s := range stories {
		storiesForTmpl[i] = CompStoryValues{
			StoryId:      s.Id(),
			StoryAuthor:  s.AuthorFUsername(),
			StoryTitle:   s.Title,
			StoryContent: s.Contents,
			StoryURL:     utils.LeftLogRight[string](url.JoinPath("/a", historyName, "story", s.Id())),
			ClampContent: true,
		}
	}

	// TODO - if not peered with this history -> no option to create story, write answers,
	ExecTemplHtmxSensitive(tmpl, w, r, "age", utils.LeftLogRight(url.JoinPath("/a", historyName, ageName)), &PageAgeValues{
		CompAgeHeaderValues: CreateAgeHeader(historyName, ageName),
		WriteStory:          CreateCompWriteStory(utils.LeftLogRight(url.JoinPath("/a", historyName, ageName, "create-story"))),
		Stories:             storiesForTmpl,
		PageBaseValues:      CreatePageBaseValues(ageName, n.Self(), history, r),
	})
}

func (n *NoUndo) HandleAgeShortcut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, utils.LeftLogRight(url.JoinPath("/a", n.Self().GetName(), mux.Vars(r)["age"])), http.StatusPermanentRedirect)
}
