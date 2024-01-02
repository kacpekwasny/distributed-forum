package noundo

import (
	"net/http"
	"strconv"

	"github.com/kacpekwasny/noundo/pkg/utils"
	"github.com/samber/mo"
)

func (n *NoUndo) HandleSignIn(w http.ResponseWriter, r *http.Request) {

}

func (n *NoUndo) HandleIndex(w http.ResponseWriter, r *http.Request) {
	self := n.Self()
	ages, err := self.GetAges(
		0,
		int(mo.TupleToResult(strconv.ParseInt(r.URL.Query().Get("ages_num"), 10, 32)).OrElse(50)),
		nil,
		nil,
	)
	utils.Loge(err)

	utils.ExecTemplLogErr(
		tplPages,
		w,
		"index.go.html",
		IndexValues{
			DisplayName: self.GetName(),
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
