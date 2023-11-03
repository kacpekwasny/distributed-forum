package forum

type LoginRequestPostBody struct {
	Login string
}

type JWTFields struct {
	Login              string
	JWTIssuedTimestamp uint64
}
