package noundo

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func (n *NoUndo) HandleHome(w http.ResponseWriter, r *http.Request) {
	self := n.Self()
	ages, err := self.GetAges(
		0,
		int(utils.GetQueryParamInt(r, "ages_num", 50)),
		nil,
		nil,
	)
	utils.Loge(err)

	ExecTemplHtmxSensitive(
		tmpl, w, r, "home", "/",
		HomeValues{
			DisplayName: self.GetName(),
			LocalAges: utils.Map(
				ages,
				func(a AgeIface) AgeLink {
					return CreateAgeInfo("/", a)
				},
			),
			Peers: utils.Map(n.Peers(), CreateHistoryInfo),
		},
	)
}

func (n *NoUndo) HandleAge(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	historyName := params["history"]
	age := params["age"]
	history, err := n.uni.GetHistoryByName(historyName)
	if err != nil {
		Handle404(w, r)
		return
	}

	storiesIface, err := history.GetStories(
		[]string{age},
		int(utils.GetQueryParamInt(r, "start", 0)),
		int(utils.GetQueryParamInt(r, "end", 50)),
		nil, nil,
	)

	if err != nil {
		// TODO, logging, user info, maybe CreateAge option?
		n.HandleHome(w, r)
		return
	}

	stories := make([]CompStoryValues, len(storiesIface))
	for i, s := range storiesIface {
		stories[i] = CompStoryValues{
			Id:              string(s.Id()),
			AuthorFUsername: s.AuthorFUsername(),
			Content:         s.Content(),
		}
	}

	ExecTemplHtmxSensitive(tmpl, w, r, "age", "/a/"+age, &PageAgeValues{
		Name:        age,
		WriteStory:  CreateCompWriteStory("/a/" + age + "/create-story"),
		Description: "TODO, description is hadrdcoded rn.",
		Stories:     stories,
	})
}
