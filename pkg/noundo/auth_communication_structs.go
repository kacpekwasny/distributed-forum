package noundo

// The basic REST Response
type RestResp struct {
	Ok      bool
	MsgCode MsgEnum
}

// Information required from User
// Information required to be held in body of http.Request
// that will be used to
type SignInRequest struct {
	Email    string
	Password string
}

type SignInResponse struct {
	RestResp
}

type SignUpRequest struct {
	Email    string
	Username string
	Password string
}

type SignUpResponse struct {
	RestResp
}

// TODO: UserFUsername field in JWT
type JWTFields struct {
	Username           string
	ParentServer       string
	JWTIssuedTimestamp int64
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
	Unauthorized    MsgEnum = "unauthorized"
	InvalidValue    MsgEnum = "invalid_value"
	InternalError   MsgEnum = "internal_error"
	InvalidURL      MsgEnum = "invalid_url"
	InvalidHeaders  MsgEnum = "invalid_headers"
	NotFound        MsgEnum = "not_found"
)
