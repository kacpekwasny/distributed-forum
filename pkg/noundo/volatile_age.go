package noundo

// a *AgeVolatile AgeIface

type AgeVolatile struct {
	name            string
	ownerUsername   string
	adminsUsernames []string
	description     string
}

func (a *AgeVolatile) GetName() string {
	return a.name
}

func (a *AgeVolatile) SetName(name string) error {
	a.name = name
	return nil
}

func (a *AgeVolatile) GetDescription() string {
	return a.description
}

func (a *AgeVolatile) GetOwner() UserIdentityIface {
	panic("TODO")
	// return a.auth.GetUserByEmail(a.ownerUsername), nil
}

func (a *AgeVolatile) ChangeOwner(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetAdmins() ([]UserPublicIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) AddAdmin(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) RemoveAdmin(user UserIdentityIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembers(start int, end int) ([]UserPublicIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembersNumber() (int, error) {
	panic("not implemented") // TODO: Implement
}

// // Create a Story written by an Author in a certain Age
// func (a *AgeVolatile) AddStory(author UserIdentityIface, age AgeIface, story StoryIface) (StoryIface, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (a *AgeVolatile) GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) []StoryIface {
// 	panic("not implemented") // TODO: Implement
// }
