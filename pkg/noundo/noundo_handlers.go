package noundo

import (
	"net/http"
	"strconv"

	"github.com/kacpekwasny/noundo/pkg/utils"
	"github.com/samber/mo"
)

func (n *NoUndo) HandleHome(w http.ResponseWriter, r *http.Request) {
	self := n.Self()
	ages, err := self.GetAges(
		0,
		int(mo.TupleToResult(strconv.ParseInt(r.URL.Query().Get("ages_num"), 10, 32)).OrElse(50)),
		nil,
		nil,
	)
	utils.Loge(err)

	ExecTemplHtmxSensitive(
		tmpl,
		w,
		r,
		"home",
		"/",
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
