package noundo

import (
	"net/url"
)

// ~~~~~~  index.go.html ~~~~~~

func CreateAgeInfo(parentUrl string, age AgeIface) AgeInfo {
	name := age.GetName()
	href, _ := url.JoinPath(parentUrl, name)
	return AgeInfo{
		DisplayName: name,
		Href:        href,
	}
}

func CreateHistoryInfo(his HistoryIface) HistoryInfo {
	name := his.GetName()
	href := his.GetUrl()
	return HistoryInfo{
		DisplayName: name,
		Href:        href,
	}
}
