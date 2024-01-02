package noundo

// When I think about this, history is an interface for a database in ints goal,
// I might make the volatile database, but in the end HistoryIface and all other interfaces must made in a way,
// That is efficient with the database.
// So my idea is, that everything that cannot be associated strictly with any Iface will be History method.

// History is the name for the whole server, that contains all Ages and All Stories
// It is an interface for storage (in case of using the local history)
// Or for a peer (in case of working with a remote history)
type HistoryIface interface {

	// domain name
	GetName() string

	// Get the URL of the History. schema + domain
	GetURL() string

	// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
	CreateAge(owner UserPublicIface, ageName string) (AgeIface, error)

	//
	GetAges(start int, end int, order OrderIface[AgeIface], filter FilterIface[AgeIface]) ([]AgeIface, error)

	//
	GetStory(id Id) (StoryIface, error)

	// Get answer from anywhere in
	GetAnswer(id Id) (AnswerIface, error)

	// GetFirst n stories ordered by different atributes, from []ages,
	GetStories(start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface], ages []AgeIface) ([]StoryIface, error)
	GetStoriesUserJoined(user UserPublicIface, start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) ([]StoryIface, error)

	GetUser(username string) (UserPublicIface, error)
	AddUser(email string, username string, password string) (UserPublicIface, error)

	// TODO:
	// GetAges that user joined,
	// GetStories first n stories of user ordered by (maybe merge with the first method???)
	// GetComments first n comments of user ordered by
}
