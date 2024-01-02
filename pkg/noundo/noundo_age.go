package noundo

type AgeIface interface {
	GetId() Id
	GetName() string
	SetName(name string) error

	GetOwner() (UserFullIface, error)
	ChangeOwner(user UserFullIface) error

	GetAdmins() ([]UserFullIface, error)
	AddAdmin(user UserFullIface) error
	RemoveAdmin(user UserFullIface) error

	GetMembers(start int, end int) ([]UserFullIface, error)
	GetMembersNumber() (int, error)

	// Create a Story written by an Author in a certain Age
	AddStory(author UserFullIface, age AgeIface, story StoryIface) (StoryIface, error)
	GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) []StoryIface
}
