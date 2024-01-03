package noundo

type BaseValues struct {
	Title            string
	MainComponentURL string
}

type SignInFormValues struct {
	Err string
}

type SignUpFormValues struct {
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
	LocalAges   []AgeInfo
	Peers       []HistoryInfo
}

type AgeInfo struct {
	DisplayName string
	Href        string
}

type HistoryInfo struct {
	DisplayName string
	Href        string
}
