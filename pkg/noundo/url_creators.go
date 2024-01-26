package noundo

import (
	"net/url"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

func AgeURL(browsingHistoryName string, ageName string) string {
	return utils.LeftLogRight(url.JoinPath("/a", browsingHistoryName, ageName))
}
