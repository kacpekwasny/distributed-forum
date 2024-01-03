package noundo

import (
	"net/url"
)

// ~~~~~~  home.go.html ~~~~~~

func CreateAgeInfo(parentURL string, age AgeIface) AgeInfo {
	name := age.GetName()
	href, _ := url.JoinPath(parentURL, name)
	return AgeInfo{
		DisplayName: name,
		Href:        href,
	}
}

func CreateHistoryInfo(his HistoryPublicIface) HistoryInfo {
	name := his.GetName()
	href := his.GetURL()
	return HistoryInfo{
		DisplayName: name,
		Href:        href,
	}
}
