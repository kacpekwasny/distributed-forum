package noundo

// When I think about this, history is an interface for a database in ints goal,
// I might make the volatile database, but in the end HistoryIface and all other interfaces must made in a way,
// That is efficient with the database.
// So my idea is, that everything that cannot be associated strictly with any Iface will be History method.

// History is the name for the whole server, that contains all Ages and All Stories
type HistoryIface interface {

	// Name to be displayed. Ex.: as the value o <a> tag.
	GetName() string

	// Get the URL of the History. Ex.: value of href atribute in an <a> tag.
	GetURL() string

	// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
	CreateAge(owner UserIface, name string) (AgeIface, error)
	GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error)

	//
	GetStory(id Id) (StoryIface, error)

	// Get answer from anywhere in
	GetAnswer(id Id) (AnswerIface, error)

	// GetFirst n stories ordered by different atributes, from []ages,
	GetStories(start int, end int, order OrderIface, filter FilterIface, ages []AgeIface) ([]StoryIface, error)

	GetUser(username string) (UserIface, error)
	AddUser(email string, username string, password string) (UserIface, error)

	// TODO:
	// GetAges that user joined,
	// GetStories first n stories of user ordered by (maybe merge with the first method???)
	// GetComments first n comments of user ordered by
}
