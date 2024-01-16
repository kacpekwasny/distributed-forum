package noundo

import (
	"net/http"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

// ~~ SignIn ~~

func (n *NoUndo) HandleSignInGet(w http.ResponseWriter, r *http.Request) {
	// todo SignIn oauth2
	jwt := GetJWTFieldsFromContext(r.Context())
	if jwt == nil {
		ExecTemplHtmxSensitive(tmpl, w, r, "signin", "signin", nil)
		return
	}
	ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", SignInFormValues{IsUser: IsUser{jwt.Username}})
}

func (n *NoUndo) HandleSignInPost(w http.ResponseWriter, r *http.Request) {
	err := SignInUser(n.Self().Authenticator(), w, r)
	if err != nil {
		ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", SignInFormValues{Err: "Sign In Failed :c"})
		return
	}
	n.HandleHome(w, r)
}

// ~~ SignUp ~~

func (n *NoUndo) HandleSignUpGet(w http.ResponseWriter, r *http.Request) {
	jwt := GetJWTFieldsFromContext(r.Context())
	if jwt == nil {
		ExecTemplHtmxSensitive(tmpl, w, r, "signup", "/signup", nil)
		return
	}
	ExecTemplHtmxSensitive(tmpl, w, r, "signup", "/signup", SignUpFormValues{IsUser: IsUser{jwt.Username}})
}

func (n *NoUndo) HandleSignUpPost(w http.ResponseWriter, r *http.Request) {
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
		ExecTemplHtmxSensitive(tmpl, w, r, "signup", "/signup", rfv)
		return
	}
	w.Header().Set("HX-Push-Url", "/signin")
	ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", utils.Ms{"Email": signUp.Email})
}

// ~~ Sign Out ~~
func (n *NoUndo) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	SignOutUser(w)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
