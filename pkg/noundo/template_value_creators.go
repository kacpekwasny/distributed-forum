package noundo

import (
	"net/url"
)

// ~~~~~~  home.go.html ~~~~~~

func CreateAgeInfo(parentDomainURL string, age AgeIface) AgeLink {
	name := age.GetName()
	href, _ := url.JoinPath(parentDomainURL, "a", name)
	return AgeLink{
		Name: name,
		Href: href,
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
