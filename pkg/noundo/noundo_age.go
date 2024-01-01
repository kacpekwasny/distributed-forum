package noundo

type AgeIface interface {
	GetId() Id
	GetName() string
	SetName(name string) error

	GetOwner() (UserIface, error)
	ChangeOwner(user UserIface) error

	GetAdmins() ([]UserIface, error)
	AddAdmin(user UserIface) error
	RemoveAdmin(user UserIface) error

	GetMembers(start int, end int) ([]UserIface, error)
	GetMembersNumber() (int, error)

	// Create a Story written by an Author in a certain Age
	AddStory(author UserIface, age AgeIface, story StoryIface) (StoryIface, error)
	GetStories(start int, end int, order OrderIface, filter FilterIface) []StoryIface
}
