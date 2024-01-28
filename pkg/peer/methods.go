package peer

// todo implement interfaces from `noundu_content_iface.go`
func (u *UserIdentity) GetFUsername() string {
	return u.Username + "@" + u.ParentServerName
}

func (u *UserPublicInfo) GetFUsername() string {
	return u.User.GetFUsername()
}

func (u *UserPublicInfo) GetParentServerName() string {
	return u.User.GetParentServerName()
}

func (u *UserPublicInfo) GetUsername() string {
	return u.User.GetUsername()
}
