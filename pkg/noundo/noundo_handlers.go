package noundo

import (
	"log/slog"
	"net/http"
	"time"

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

func (n *NoUndo) HandleSelfProfile(w http.ResponseWriter, r *http.Request) {
	userJWT := GetJWT(r.Context())
	if userJWT == nil {
		http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		return
	}
	user, err := n.Self().GetUser(userJWT.Username())
	if err != nil {
		slog.Error("cannot retrieve user from database, but has valid JWT", "username", userJWT.Username, "parent server", userJWT.parentServerName)
		utils.WriteJsonWithStatus(w, "my apologies, you don't exist", http.StatusInternalServerError)
		return
	}
	ExecTemplHtmxSensitive(tmpl, w, r, "profile", "/profile", PageProfileValues{
		Username:         user.Username(),
		ParentServerName: user.ParentServerName(),
		AccountBirthDate: time.Unix(user.AccountBirthDate(), 0).Format(time.RFC3339),
		AboutMe:          user.AboutMe(),
		SelfProfile:      true,
		PageBaseValues:   CreatePageBaseValues("My Profile", n.Self(), n.Self(), r),
	})
}

func (n *NoUndo) HandleProfile(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := n.Self().GetUser(username)
	if err != nil {
		n.Handle404(w, r)
		return
	}

	// TODO Jwt handle more information and then change JoinURL for ProfileURL
	ExecTemplHtmxSensitive(tmpl, w, r, "profile", JoinURL("/profile", username), PageProfileValues{
		Username:         user.Username(),
		ParentServerName: "@" + user.ParentServerName(),
		AccountBirthDate: "todo birthdate",
		AboutMe:          "todo - keep user aboutme - only editable thing",
		SelfProfile:      false,
		PageBaseValues:   CreatePageBaseValues("Profile", n.Self(), n.Self(), r),
	})
}
