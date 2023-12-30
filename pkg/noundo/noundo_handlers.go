package noundo

import (
	"net/http"
	"strconv"

	"github.com/kacpekwasny/noundo/pkg/auth"
	"github.com/kacpekwasny/noundo/pkg/utils"
	"github.com/samber/mo"
)

func (n *NoUndo) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (n *NoUndo) handleIndex(w http.ResponseWriter, r *http.Request) {
	jwtf := auth.GetJWTFieldsFromContext(r.Context())
	if jwtf == nil {
		// TODO - proper template
		utils.ExecTemplLogErr(tplPages, w, "welcome.go.html", WelcomeValues{Msg: "It's a shame you didn't Log In :(("})
		return
	}
	self := n.Self()

	ages, err := self.GetAges(
		0,
		int(mo.TupleToResult(strconv.ParseInt(r.URL.Query().Get("agesnum"), 10, 32)).OrElse(50)),
		nil,
		nil,
	)
	utils.Loge(err)

	utils.ExecTemplLogErr(
		tplPages,
		w,
		"welcome.go.html",
		IndexValues{
			DisplayName: self.GetURL(),
			LocalAges: utils.Map(
				ages,
				func(a AgeIface) AgeInfo {
					return CreateAgeInfo(self.GetURL(), a)
				},
			),
			Peers: utils.Map(n.Peers(), CreateHistoryInfo),
		},
	)
}
