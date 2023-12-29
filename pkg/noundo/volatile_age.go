package noundo

import "github.com/kacpekwasny/distributed-forum/pkg/auth"

// a *AgeVolatile AgeIface

type AgeVolatile struct {
	id              Id
	name            string
	ownerUsername   string
	adminsUsernames []string
	auth            auth.Authenticator
}

func (a *AgeVolatile) GetId() Id {
	return a.id
}

func (a *AgeVolatile) GetName() string {
	return a.name
}

func (a *AgeVolatile) SetName(name string) error {
	a.name = name
	return nil
}

func (a *AgeVolatile) GetOwner() (UserIface, error) {
	panic("TODO")
	// return a.auth.GetUserByLogin(a.ownerUsername), nil
}

func (a *AgeVolatile) ChangeOwner(user UserIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetAdmins() ([]UserIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) AddAdmin(user UserIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) RemoveAdmin(user UserIface) error {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembers(start int, end int) ([]UserIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembersNumber() (int, error) {
	panic("not implemented") // TODO: Implement
}

// Create a Story written by an Author in a certain Age
func (a *AgeVolatile) AddStory(author UserIface, age AgeIface, story StoryIface) (StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetStories(start int, end int, order OrderIface, filter FilterIface) []StoryIface {
	panic("not implemented") // TODO: Implement
}
