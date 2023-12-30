package noundo

// The basic REST Response
type RestResp struct {
	Ok      bool
	MsgCode MsgEnum
}

// Information required from User
// Information required to be held in body of http.Request
// that will be used to
type LoginMe struct {
	Email    string
	Password string
}

type LoginMeResponse struct {
	RestResp
}

type RegisterMe struct {
	Email    string
	Username string
	Password string
}

type RegisterMeResponse struct {
	RestResp
}

// TODO: UserId field in JWT
type JWTFields struct {
	Username           string
	JWTIssuedTimestamp uint64
}

type MsgEnum string

const (
	Ok              MsgEnum = "ok"
	Err             MsgEnum = "err"
	DecodeErr       MsgEnum = "decode_err"
	TokenSignErr    MsgEnum = "token_signed_err"
	EmailInUse      MsgEnum = "email_in_use"
	UsernameInUse   MsgEnum = "username_in_use"
	InvalidPassword MsgEnum = "invalid_pass"
)
