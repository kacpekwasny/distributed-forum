package noundo

// The basic REST Response
type RestResp struct {
	Ok      bool
	MsgCode MsgCode
}

// Information required from User
// Information required to be held in body of http.Request
// that will be used to
type LoginMe struct {
	Login    string
	Password string
}

type LoginMeResponse struct {
	RestResp
}

type RegisterMe struct {
	Login    string
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

type MsgCode string

const (
	Ok              MsgCode = "ok"
	Err             MsgCode = "err"
	DecodeErr       MsgCode = "decode_err"
	TokenSignErr    MsgCode = "token_signed_err"
	LoginInUse      MsgCode = "login_in_use"
	UsernameInUser  MsgCode = "username_in_use"
	InvalidPassword MsgCode = "invalid_pass"
)
