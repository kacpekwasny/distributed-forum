package noundo

type BaseValues struct {
	Title            string
	MainComponentURL string
	Navbar           NavbarValues
}

type NavbarValues struct {
	UsingHistoryName    string
	BrowsingHistoryName string
	BrowsingHistoryURL  string
	UserProfile         bool
}

type IsUser struct {
	CurrentUsername string
}

type SignInFormValues struct {
	IsUser
	Err string
}

type SignUpFormValues struct {
	IsUser
	Email    string
	Username string

	ErrEmail    string
	ErrUsername string
	ErrPassword string

	Err string
}

type WelcomeValues struct {
	Username string
	Msg      string
}

// ~~~~~~  home.go.html ~~~~~~

type HomeValues struct {
	DisplayName string
	LocalAges   []AgeLink
	Peers       []HistoryInfo
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
	NavbarValues
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
