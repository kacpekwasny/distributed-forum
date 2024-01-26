package noundo

import (
	"net/url"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

func AgeURL(browsingHistoryName string, ageName string) string {
	return utils.LeftLogRight(url.JoinPath("/a", browsingHistoryName, ageName))
}

func StoryURL(browsingHistoryName string, storyId string) string {
	return utils.LeftLogRight(url.JoinPath("/a", browsingHistoryName, "story", storyId))
}

func ProfileURL(user UserPublicIface, usingHistoryName string) string {
	if usingHistoryName == user.ParentServerName() {
		return utils.LeftLogRight(url.JoinPath("/profile", user.Username()))
	}
	return utils.LeftLogRight(url.JoinPath("https://"+user.ParentServerName(), "profile", user.Username()))
}
