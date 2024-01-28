package noundo

// a *AgeVolatile AgeIface

type AgePersistent struct {
	id          string
	name        string
	ownerId     string
	description string
}

func (a *AgePersistent) GetId() string {
	return a.id
}

func (a *AgePersistent) GetName() string {
	return a.name
}

func (a *AgePersistent) SetName(name string) error {
	a.name = name
	return nil
}

func (a *AgePersistent) GetDescription() string {
	return a.description
}

func (a *AgePersistent) GetOwner() (UserPublicIface, error) {
	panic("TODO")
	// return a.auth.GetUserByEmail(a.ownerUsername), nil
}

func (a *AgePersistent) ChangeOwner(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgePersistent) GetAdmins() ([]UserPublicIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgePersistent) AddAdmin(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgePersistent) RemoveAdmin(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgePersistent) GetMembers(start int, end int) ([]UserPublicIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgePersistent) GetMembersNumber() (int, error) {
	panic("not implemented") // TODO: Implement
}

// // Create a Story written by an Author in a certain Age
// func (a *AgeVolatile) AddStory(author UserIdentityIface, age AgeIface, story StoryIface) (StoryIface, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (a *AgeVolatile) GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) []StoryIface {
// 	panic("not implemented") // TODO: Implement
// }
