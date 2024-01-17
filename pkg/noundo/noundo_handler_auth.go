package noundo

import (
	"net/http"
)

// ~~ SignIn ~~

func (n *NoUndo) HandleSignInGet(w http.ResponseWriter, r *http.Request) {
	// todo SignIn oauth2
	ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", PageSignInValues{
		PageBaseValues: CreatePageBaseValues("Sign In", n.Self(), n.Self(), r),
	})
}

func (n *NoUndo) HandleSignInPost(w http.ResponseWriter, r *http.Request) {
	err := SignInUser(n.Self().Authenticator(), w, r)
	if err != nil {
		ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", PageSignInValues{
			PageBaseValues: CreatePageBaseValues("Sign In", n.Self(), n.Self(), r),
			Err:            "Sign In Failed :c",
		})
		return
	}
	// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	n.HandleHome(w, r)
}

// ~~ SignUp ~~

func (n *NoUndo) HandleSignUpGet(w http.ResponseWriter, r *http.Request) {
	ExecTemplHtmxSensitive(tmpl, w, r, "signup", "/signup", PageSignUpValues{
		PageBaseValues: CreatePageBaseValues("Sign In", n.Self(), n.Self(), r),
	})
}

func (n *NoUndo) HandleSignUpPost(w http.ResponseWriter, r *http.Request) {
	signUp, err := GetSignUpRequest(r)

	var regResp *SignUpResponse
	if err != nil {
		regResp = &SignUpResponse{RestResp: RestResp{Ok: false, MsgCode: DecodeErr}}
	} else {
		regResp = n.uni.Authenticator().SignUpUser(signUp)
	}

	var rfv PageSignUpValues

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
	ExecTemplHtmxSensitive(tmpl, w, r, "signin", "/signin", PageSignInValues{
		PageBaseValues: CreatePageBaseValues("Sign In", n.Self(), n.Self(), r),
		Email:          signUp.Email,
	})
}

// ~~ Sign Out ~~
func (n *NoUndo) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	SignOutUser(w)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
