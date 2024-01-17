package noundo

type PageBaseValues struct {
	Title string
	CompNavbarValues
	UserInfo
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

type SignUpFormValues struct {
	UserInfo
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
	Name        string
	WriteStory  CompWriteStory
	Description string
	Stories     []CompStoryValues
	PageBaseValues
}

type CompWriteStory struct {
	HxPost        string
	TitleLenMin   int
	TitleLenMax   int
	ContentLenMin int
	ContentLenMax int
}

type CompStoryValues struct {
	Id              string
	AuthorFUsername string
	Content         string

	// TODO answers
}

type PageProfileValues struct {
	Username         string
	ParentServerName string
	AccountBirthDate string
	AboutMe          string
	SelfProfile      bool
	PageBaseValues
}
