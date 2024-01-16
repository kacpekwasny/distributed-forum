package noundo

import (
	"net/http"
	"net/url"
)

// ~~~~~~  home.go.html ~~~~~~

func createAgeInfo(parentDomainURL string, historyName string, ageName string) AgeLink {
	href, _ := url.JoinPath(parentDomainURL, "a", historyName, ageName)
	return AgeLink{
		Name: ageName,
		Href: href,
	}
}

func createHistoryInfo(his HistoryPublicIface) HistoryInfo {
	name := his.GetName()
	href := his.GetURL()
	return HistoryInfo{
		DisplayName: name,
		Href:        href,
	}
}

func createCompWriteStory(hxPost string) CompWriteStory {
	return CompWriteStory{
		HxPost:        hxPost,
		TitleLenMin:   TITLE_LEN_MIN,
		TitleLenMax:   TITLE_LEN_MAX,
		ContentLenMin: CONTENT_LEN_MIN,
		ContentLenMax: CONTENT_LEN_MAX,
	}
}

func createPageBaseValues(title string, using HistoryPublicIface, browsing HistoryPublicIface, r *http.Request) PageBaseValues {
	return PageBaseValues{
		Title: title,
		CompNavbarValues: CompNavbarValues{
			UsingHistoryName:    using.GetName(),
			BrowsingHistoryName: browsing.GetName(),
			BrowsingHistoryURL:  browsing.GetURL(),
			SignedIn:            jwtInCtx(r),
		},
	}
}
