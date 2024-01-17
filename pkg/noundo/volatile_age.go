package noundo

// a *AgeVolatile AgeIface

type AgeVolatile struct {
	id              string
	name            string
	ownerUsername   string
	adminsUsernames []string
	auth            AuthenticatorIface
}

func (a *AgeVolatile) GetId() string {
	return a.id
}

func (a *AgeVolatile) GetName() string {
	return a.name
}

func (a *AgeVolatile) SetName(name string) error {
	a.name = name
	return nil
}

func (a *AgeVolatile) GetOwner() (UserFullIface, error) {
	panic("TODO")
	// return a.auth.GetUserByEmail(a.ownerUsername), nil
}

func (a *AgeVolatile) ChangeOwner(user UserFullIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetAdmins() ([]UserFullIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) AddAdmin(user UserFullIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) RemoveAdmin(user UserFullIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembers(start int, end int) ([]UserFullIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembersNumber() (int, error) {
	panic("not implemented") // TODO: Implement
}

// Create a Story written by an Author in a certain Age
func (a *AgeVolatile) AddStory(author UserFullIface, age AgeIface, story StoryIface) (StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) []StoryIface {
	panic("not implemented") // TODO: Implement
}
