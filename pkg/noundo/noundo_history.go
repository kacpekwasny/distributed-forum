package noundo

// TODO - filter posts by hot, best, most reactions, new, administered
// Admin may block post from appearing in feed apart of sorted by `NEW` and `most reactions`
// Admin may change the Age's description,

// When I think about this, history is an interface for a database in ints goal,
// I might make the volatile database, but in the end HistoryIface and all other interfaces must made in a way,
// That is efficient with the database.
// So my idea is, that everything that cannot be associated strictly with any Iface will be History method.

// History is the name for the whole server, that contains all Ages and All Stories
// It is an interface for storage (in case of using the local history)
// Or for a peer (in case of working with a remote history)
type HistoryReadIface interface {
	// domain name
	GetName() string
	// Get the URL of the History. schema + domain
	GetURL() string

	// Retrive all user info by supplying a username
	GetUser(username string) (UserPublicIface, error)
	GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error)

	// Get a single story
	GetStory(id string) (Story, error)
	// Get `n` stories ordered by different atributes, from []ages,
	GetStories(ageNames []string, start int, end int, order OrderIface, filter FilterIface) ([]*Story, error)

	// Get answer from anywhere in
	GetAnswer(id string) (Answer, error)
	// Get tree of answers, to the specified postable with the specified depth
	GetAnswers(postableId string, start int, end int, depth int, order OrderIface, filter FilterIface) ([]*Story, error)

	// todo later
	// GetAgeOwner(name string) (UserIdentityIface, error)
	// GetAgeAdmins(name string) ([]UserIdentityIface, error)
	// GetAgeMembers(start int, end int) ([]UserPublicIface, error)
	// GetAgeMembersNumber(name string) (int, error)
}

type HistoryWriteIface interface {
	// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
	CreateAge(owner UserIdentityIface, ageName string) (AgeIface, error)
	// Create a new story within
	CreateStory(author UserIdentityIface, ageName string, story StoryContent) (Story, error)
	// Create an Answer under a post or other Answer
	CreateAnswer(author UserIdentityIface, parentId string, answerContent string) (Answer, error)
}

type HistoryWritePrivateIface interface {
	Authenticator() AuthenticatorIface

	// Create a new user
	CreateUser(email string, username string, password string) (UserPublicIface, error)
}

type HistoryPublicIface interface {
	HistoryReadIface
	HistoryWriteIface
}

type HistoryFullIface interface {
	HistoryReadIface
	HistoryWriteIface
	HistoryWritePrivateIface
}
