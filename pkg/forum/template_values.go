package forum

type BaseValues struct {
	Title          string
	MainContentUrl string
}

type IndexValues struct {
	Text1 string
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
