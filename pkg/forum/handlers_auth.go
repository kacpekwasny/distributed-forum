package forum

import (
	"net/http"

	"github.com/kacpekwasny/distributed-forum/pkg/auth"
	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

func BaseGetFactory(baseValues BaseValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "base.go.html", baseValues)
		utils.Pife(err)
	}
}

func ComponentGetFactory(template string, v any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, template, v)
		utils.Pife(err)
	}
}

// ~~ Login ~~

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	err := auth.LoginUser(authenticator, w, r)
	if err != nil {
		err = tpl.ExecuteTemplate(w, "login.go.html", LoginFormValues{Err: "Login Failed :c"})
		utils.Pife(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/")
	err = tpl.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "Welcome :D",
			MainContentUrl: "welcome"},
	)
	utils.Pife(err)
}

// ~~ Register ~~

func HandlePostRegister(w http.ResponseWriter, r *http.Request) {
	regResp := auth.RegisterUser(authenticator, r)

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
		err := tpl.ExecuteTemplate(w, "register.go.html", rfv)
		utils.Pife(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/login")
	err := tpl.ExecuteTemplate(w, "login.go.html", rfv)
	utils.Pife(err)
}
