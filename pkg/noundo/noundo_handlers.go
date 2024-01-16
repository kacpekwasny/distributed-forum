package noundo

import (
	"log/slog"
	"net/http"
	"net/url"

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
		PageHomeValues{
			DisplayName: self.GetName(),
			LocalAges: utils.Map(
				ages,
				func(a AgeIface) AgeLink {
					return CreateAgeInfo("/", n.Self().GetName(), a.GetName())
				},
			),
			Peers:          utils.Map(n.Peers(), CreateHistoryInfo),
			PageBaseValues: CreatePageBaseValues(n.Self().GetName(), n.Self(), n.Self(), r),
		},
	)
}

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

	storiesIface, err := history.GetStories(
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

	stories := make([]CompStoryValues, len(storiesIface))
	for i, s := range storiesIface {
		stories[i] = CompStoryValues{
			Id:              string(s.Id()),
			AuthorFUsername: s.AuthorFUsername(),
			Content:         s.Content(),
		}
	}

	// TODO - if not peered with this history -> no option to create story, write answers,
	ExecTemplHtmxSensitive(tmpl, w, r, "age", utils.LeftLogRight(url.JoinPath("/a", historyName, ageName)), &PageAgeValues{
		Name:           ageName,
		WriteStory:     CreateCompWriteStory(utils.LeftLogRight(url.JoinPath("/a", historyName, ageName, "create-story"))),
		Description:    "TODO, description is hadrdcoded rn.",
		Stories:        stories,
		PageBaseValues: CreatePageBaseValues(ageName, n.Self(), history, r),
	})
}

func (n *NoUndo) HandleAgeShortcut(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, utils.LeftLogRight(url.JoinPath("/a", n.Self().GetName(), mux.Vars(r)["age"])), http.StatusPermanentRedirect)
}

func (n *NoUndo) HandleSelfProfile(w http.ResponseWriter, r *http.Request) {
	userJWT := GetJWTFieldsFromContext(r.Context())
	if userJWT == nil {
		http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		return
	}
	user, err := n.Self().GetUser(userJWT.Username)
	if err != nil {
		slog.Error("cannot get user, but is logged in", userJWT.Username, userJWT.UserFUsername)
	}
	ExecTemplHtmxSensitive(tmpl, w, r, "profile", "/profile", PageProfileValues{
		Username:         userJWT.Username,
		ParentServerName: "@" + user.ParentServerName(),
		AccountBirthDate: "todo birthdate",
		AboutMe:          "todo - keep user aboutme - only editable thing",
		SelfProfile:      true,
		PageBaseValues:   CreatePageBaseValues("Title", n.Self(), n.Self(), r),
	})
}

func (n *NoUndo) HandleProfile(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := n.Self().GetUser(username)
	if err != nil {
		n.Handle404(w, r)
		return
	}

	ExecTemplHtmxSensitive(tmpl, w, r, "profile", utils.LeftLogRight(url.JoinPath("/profile", username)), PageProfileValues{
		Username:         user.Username(),
		ParentServerName: "@" + user.ParentServerName(),
		AccountBirthDate: "todo birthdate",
		AboutMe:          "todo - keep user aboutme - only editable thing",
		SelfProfile:      false,
		PageBaseValues:   CreatePageBaseValues("Profile", n.Self(), n.Self(), r),
	})
}
