package noundo

import (
	"net/url"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

func AgeURL(browsingHistoryName string, ageName string) string {
	return JoinURL("/a", browsingHistoryName, ageName)
}

func StoryURL(browsingHistoryName string, storyId string) string {
	return JoinURL("/a", browsingHistoryName, "story", storyId)
}

func ProfileURL(user UserIdentityIface, usingHistoryName string) string {
	if usingHistoryName == user.GetParentServerName() {
		return JoinURL("/profile", user.GetUsername())
	}
	return JoinURL("https://"+user.GetParentServerName(), "profile", user.GetUsername())
}

func WriteStoryURL(historyName string, ageName string) string {
	return JoinURL("/a", historyName, ageName, "create-story")
}

func WriteAnswerURL(historyName string, storyId string) string {
	return JoinURL("/write-answer", historyName, storyId)
}
func JoinURL(base string, elem ...string) string {
	return utils.LeftLogRight(url.JoinPath(base, elem...))
}
