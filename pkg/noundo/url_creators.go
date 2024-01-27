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

func ProfileURL(user UserPublicIface, usingHistoryName string) string {
	if usingHistoryName == user.ParentServerName() {
		return JoinURL("/profile", user.Username())
	}
	return JoinURL("https://"+user.ParentServerName(), "profile", user.Username())
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
