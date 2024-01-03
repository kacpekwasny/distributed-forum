package noundo

import (
	"net/http"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

// ~~ SignIn ~~

func (n *NoUndo) HandlePostSignIn(w http.ResponseWriter, r *http.Request) {
	err := SignInUser(n.Self().Authenticator(), w, r)
	if err != nil {
		ExecTemplHtmxSensitive(tmpl, w, r, SignInFormValues{Err: "Sign In Failed :c"}, "signin")
		return
	}
	w.Header().Set("HX-Push-Url", "/")
	ExecTemplHtmxSensitive(tmpl, w, r,
		BaseValues{
			Title:            "Welcome :D",
			MainComponentURL: "home"},
		"base",
	)
}

// ~~ SignUp ~~

func (n *NoUndo) HandlePostSignUp(w http.ResponseWriter, r *http.Request) {
	signUp, err := GetSignUpRequest(r)

	var regResp *SignUpResponse
	if err != nil {
		regResp = &SignUpResponse{RestResp: RestResp{Ok: false, MsgCode: DecodeErr}}
	} else {
		regResp = n.uni.Authenticator().SignUpUser(signUp)
	}

	var rfv SignUpFormValues

	if !regResp.Ok {
		switch regResp.MsgCode {
		case EmailInUse:
			rfv.ErrEmail = "Email is in use."
		case UsernameInUse:
			rfv.ErrUsername = "Username is in use."
		case InvalidPassword:
			rfv.ErrPassword = "Password does not match criteria."
		default:
			rfv.Err = "Unknown error occured."
		}
		ExecTemplHtmxSensitive(tmpl, w, r, rfv, "signup")
		return
	}
	w.Header().Set("HX-Push-Url", "/signin")
	ExecTemplHtmxSensitive(tmpl, w, r, utils.Ms{"Email": signUp.Email}, "signin")
}

// ~~ Sign Out ~~
func (n *NoUndo) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	SignOutUser(w)
	ExecTemplHtmxSensitive(tmpl, w, r,
		BaseValues{
			Title:            "Home",
			MainComponentURL: "home",
		},
		"base",
	)
}
