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

func CreateCompWriteStory(hxPost string) CompWriteStory {
	return CompWriteStory{
		HxPost:        hxPost,
		TitleLenMin:   TITLE_LEN_MIN,
		TitleLenMax:   TITLE_LEN_MAX,
		ContentLenMin: CONTENT_LEN_MIN,
		ContentLenMax: CONTENT_LEN_MAX,
	}
}
