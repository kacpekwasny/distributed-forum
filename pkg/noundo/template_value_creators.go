package noundo

import (
	"net/http"
	"net/url"
)

// ~~~~~~  home.go.html ~~~~~~

func CreateAgeInfo(parentDomainURL string, historyName string, ageName string) AgeLink {
	href, _ := url.JoinPath(parentDomainURL, "a", historyName, ageName)
	return AgeLink{
		Name: ageName,
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
		ContentLenMin: STORY_LEN_MIN,
		ContentLenMax: STORY_LEN_MAX,
	}
}

func CreatePageBaseValues(title string, using HistoryPublicIface, browsing HistoryPublicIface, r *http.Request) PageBaseValues {
	return PageBaseValues{
		PageTitle:       title,
		CurrentUserInfo: CreateCurrentUserInfo(r),
		CompNavbarValues: CompNavbarValues{
			UsingHistoryName:    using.GetName(),
			BrowsingHistoryName: browsing.GetName(),
			BrowsingHistoryURL:  browsing.GetURL(),
		},
	}
}

func CreateCurrentUserInfo(r *http.Request) CurrentUserInfo {
	jwt := GetJWT(r.Context())
	if jwt == nil {
		return CurrentUserInfo{
			SignedIn: false,
		}
	}
	return CurrentUserInfo{
		Username: jwt.Username,
		SignedIn: true,
	}
}

func CreateAgeHeader(browsingHistoryName string, ageName string) CompAgeHeaderValues {
	return CompAgeHeaderValues{
		AgeName:     ageName,
		AgeURL:      AgeURL(browsingHistoryName, ageName),
		Description: "todo description",
	}
}
