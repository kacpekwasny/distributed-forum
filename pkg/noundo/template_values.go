package noundo

type BaseValues struct {
	Title            string
	MainComponentURL string
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
	Description string
	Stories     []CompStoryValues
}

type CompStoryValues struct {
	Id              string
	AuthorFUsername string
	Content         string

	// TODO answers
}
