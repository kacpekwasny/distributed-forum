package noundo

type BaseValues struct {
	Title          string
	MainContentUrl string
}

type LoginFormValues struct {
	Err string
}

type RegisterFormValues struct {
	Login    string
	Username string

	ErrLogin    string
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
