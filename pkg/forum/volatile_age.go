package forum

import "github.com/kacpekwasny/distributed-forum/pkg/auth"

// a *AgeVolatile AgeIface

type AgeVolatile struct {
	id              Id
	name            string
	ownerUsername   string
	adminsUsernames []string
	auth            auth.Authenticator
}

func (a *AgeVolatile) GetId() (Id, error) {
	return a.id, nil
}

func (a *AgeVolatile) GetName() (string, error) {
	return a.name, nil
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

func (a *AgeVolatile) GetMembers(start uint32, end uint32) ([]UserIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetMembersNumber() (uint32, error) {
	panic("not implemented") // TODO: Implement
}

// Create a Story written by an Author in a certain Age
func (a *AgeVolatile) AddStory(author UserIface, age AgeIface, story StoryIface) (StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AgeVolatile) GetStories(start uint32, end uint32, order OrderIface, filter FilterIface) []StoryIface {
	panic("not implemented") // TODO: Implement
}
