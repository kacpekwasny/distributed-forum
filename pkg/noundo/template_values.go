package noundo

type BaseValues struct {
	Title          string
	MainContentURL string
}

type LoginFormValues struct {
	Err string
}

type RegisterFormValues struct {
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

// ~~~~~~  index.go.html ~~~~~~

type IndexValues struct {
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
