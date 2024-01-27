package noundo

import (
	"net/http"
	"net/url"

	"github.com/kacpekwasny/noundo/pkg/utils"
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
	ages, _ := his.GetAges(0, 30, nil, nil)
	return HistoryInfo{
		DisplayName: name,
		Href:        href,
		Ages: utils.Map(ages, func(a AgeIface) AgeLink {
			return AgeLink{
				Name: a.GetName(),
				Href: AgeURL(his.GetName(), a.GetName()),
			}
		}),
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
		Username: jwt.Username(),
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
