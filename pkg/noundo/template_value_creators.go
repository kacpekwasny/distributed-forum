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
		ContentLenMin: CONTENT_LEN_MIN,
		ContentLenMax: CONTENT_LEN_MAX,
	}
}

func CreatePageBaseValues(title string, using HistoryPublicIface, browsing HistoryPublicIface, r *http.Request) PageBaseValues {
	return PageBaseValues{
		Title:    title,
		UserInfo: CreateUserInfo(r),
		CompNavbarValues: CompNavbarValues{
			UsingHistoryName:    using.GetName(),
			BrowsingHistoryName: browsing.GetName(),
			BrowsingHistoryURL:  browsing.GetURL(),
			SignedIn:            jwtInCtx(r),
		},
	}
}

func CreateUserInfo(r *http.Request) UserInfo {
	jwt := GetJWTFieldsFromContext(r.Context())
	if jwt == nil {
		return UserInfo{
			SignedIn: false,
		}
	}
	return UserInfo{
		Username: jwt.Username,
		SignedIn: true,
	}
}
