package peer

// todo implement interfaces from `noundu_content_iface.go`
func (u *UserIdentity) GetFUsername() string {
	if u != nil {
		return u.Username + "@" + u.ParentServerName
	}
	return ""
}
