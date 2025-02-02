package noundo

type AgeIface interface {
	GetName() string
	GetOwner() UserIdentityIface
	GetDescription() string

	// GetAdmins() ([]UserPublicIface, error)
	// AddAdmin(user UserIdentityIface) error
	// RemoveAdmin(user UserIdentityIface) error

	// GetMembers(start int, end int) ([]UserPublicIface, error)
	// GetMembersNumber() (int, error)

	// Create a Story written by an Author in a certain Age
	// AddStory(author UserIdentityIface, age AgeIface, story StoryIface) (StoryIface, error)
	// GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) []StoryIface
}
