package noundo

type PageBaseValues struct {
	CompNavbarValues
	UserInfo

	PageTitle string
}

type UserInfo struct {
	Username string
	SignedIn bool
}

type CompNavbarValues struct {
	UsingHistoryName    string
	BrowsingHistoryName string
	BrowsingHistoryURL  string
}

type PageSignInValues struct {
	PageBaseValues

	Email string
	Err   string
}

type PageSignUpValues struct {
	PageBaseValues

	Email    string
	Username string

	ErrEmail    string
	ErrUsername string
	ErrPassword string

	Err string
}

// ~~~~~~  home.go.html ~~~~~~

type PageHomeValues struct {
	DisplayName string
	LocalAges   []AgeLink
	Peers       []HistoryInfo
	PageBaseValues
}

type AgeLink struct {
	Name string
	Href string
}

type HistoryInfo struct {
	DisplayName string
	Href        string
}

type PageAgeValues struct {
	PageBaseValues
	CompAgeHeaderValues

	WriteStory CompWriteStory
	Stories    []CompStoryValues
}

type CompAgeHeaderValues struct {
	AgeName     string
	AgeURL      string
	Description string
}

type CompWriteStory struct {
	HxPost        string
	TitleLenMin   int
	TitleLenMax   int
	ContentLenMin int
	ContentLenMax int
}

type CompStoryValues struct {
	StoryId      string
	StoryTitle   string
	StoryAuthor  string
	StoryContent string
	StoryURL     string

	// TODO answers
}

type PageStoryValues struct {
	PageBaseValues
	CompStoryValues
	CompAgeHeaderValues
}

type PageProfileValues struct {
	PageBaseValues

	Username         string
	ParentServerName string
	AccountBirthDate string
	AboutMe          string
	SelfProfile      bool
}

type Page401Values struct {
	RequestedPath string
	PageBaseValues
}
