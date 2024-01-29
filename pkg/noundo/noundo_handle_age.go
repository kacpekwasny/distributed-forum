package noundo

import (
	"log/slog"
	"net/http"
	"strconv"

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

	stories, err := history.GetStories(
		[]string{ageName},
		int(utils.GetQueryParamInt(r, "start", 0)),
		int(utils.GetQueryParamInt(r, "end", 50)),
		nil, nil,
	)

	if err != nil {
		slog.Debug("GetStories failed", "ageName", ageName, "err", err)
		n.Handle404(w, r)
		return
	}

	storiesForTmpl := make([]CompStoryValues, len(stories))
	for i, s := range stories {
		storiesForTmpl[i] = CompStoryValues{
			Story:        s,
			ClampContent: true,
			StoryURL:     StoryURL(s., strconv.Itoa(s.Id())),
		}
	}

	// TODO - if not peered with this history -> no option to create story, write answers,
	ExecTemplHtmxSensitive(tmpl, w, r, "age", AgeURL(historyName, ageName), &PageAgeValues{
		CompAgeHeaderValues: CreateAgeHeader(historyName, ageName),
		WriteStory:          CreateCompWriteStory(WriteStoryURL(historyName, ageName)),
		Stories:             storiesForTmpl,
		PageBaseValues:      CreatePageBaseValues(ageName, n.Self(), history, r),
	})
}

func (n *NoUndo) HandleAgeShortcut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, AgeURL(n.Self().GetName(), mux.Vars(r)["age"]), http.StatusPermanentRedirect)
}
