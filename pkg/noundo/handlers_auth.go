package noundo

import (
	"net/http"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

func BaseGetFactory(baseValues BaseValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tplPages.ExecuteTemplate(w, "base.go.html", baseValues)
		utils.Loge(err)
	}
}

func ComponentGetFactory(template string, v any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tplPages.ExecuteTemplate(w, template, v)
		utils.Loge(err)
	}
}

// ~~ SignIn ~~

func HandlePostSignIn(w http.ResponseWriter, r *http.Request) {
	err := SignInUser(authenticator, w, r)
	if err != nil {
		err = tplPages.ExecuteTemplate(w, "signin.go.html", SignInFormValues{Err: "Sign In Failed :c"})
		utils.Loge(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/")
	err = tplPages.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "Welcome :D",
			MainContentURL: "welcome"},
	)
	utils.Loge(err)
}

// ~~ SignUp ~~

func HandlePostSignUp(w http.ResponseWriter, r *http.Request) {
	signUp, err := GetSignUpRequest(r)

	var regResp *SignUpResponse
	if err != nil {
		regResp = &SignUpResponse{RestResp: RestResp{Ok: false, MsgCode: DecodeErr}}
	} else {
		regResp = authenticator.SignUpUser(signUp)
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
		err := tplPages.ExecuteTemplate(w, "signup.go.html", rfv)
		utils.Loge(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/signin")
	err = tplPages.ExecuteTemplate(w, "signin.go.html", utils.Ms{"Email": signUp.Email})
	utils.Loge(err)
}

// ~~ Sign Out ~~
func HandleSignOut(w http.ResponseWriter, r *http.Request) {
	SignOutUser(w)
	utils.Loge(tplPages.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "Sign In",
			MainContentURL: "signin",
		}))
}
