package noundo

import (
	"net/http"

	"github.com/kacpekwasny/noundo/pkg/auth"
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

// ~~ Login ~~

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	err := auth.LoginUser(authenticator, w, r)
	if err != nil {
		err = tplPages.ExecuteTemplate(w, "login.go.html", LoginFormValues{Err: "Login Failed :c"})
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

// ~~ Register ~~

func HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	registerMe, err := auth.GetRegisterMe(r)

	var regResp *auth.RegisterMeResponse
	if err != nil {
		regResp = &auth.RegisterMeResponse{RestResp: auth.RestResp{Ok: false, MsgCode: auth.DecodeErr}}
	} else {
		regResp = authenticator.RegisterUser(registerMe)
	}

	var rfv RegisterFormValues

	if !regResp.Ok {
		switch regResp.MsgCode {
		case auth.LoginInUse:
			rfv.ErrLogin = "Login is in use."
		case auth.UsernameInUser:
			rfv.ErrUsername = "Username is in use."
		case auth.InvalidPassword:
			rfv.ErrPassword = "Password does not match criteria."
		default:
			rfv.Err = "Unknown error occured."
		}
		err := tplPages.ExecuteTemplate(w, "register.go.html", rfv)
		utils.Loge(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/login")
	err = tplPages.ExecuteTemplate(w, "login.go.html", utils.Ms{"Login": registerMe.Login})
	utils.Loge(err)
}

// ~~ Log Out ~~
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	auth.LogoutUser(w)
	utils.Loge(tplPages.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "Login",
			MainContentURL: "login",
		}))
}
