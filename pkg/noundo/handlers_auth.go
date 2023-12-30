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

// ~~ Login ~~

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	err := LoginUser(authenticator, w, r)
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
	registerMe, err := GetRegisterMe(r)

	var regResp *RegisterMeResponse
	if err != nil {
		regResp = &RegisterMeResponse{RestResp: RestResp{Ok: false, MsgCode: DecodeErr}}
	} else {
		regResp = authenticator.RegisterUser(registerMe)
	}

	var rfv RegisterFormValues

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
		err := tplPages.ExecuteTemplate(w, "register.go.html", rfv)
		utils.Loge(err)
		return
	}
	w.Header().Set("HX-Push-Url", "/login")
	err = tplPages.ExecuteTemplate(w, "login.go.html", utils.Ms{"Email": registerMe.Email})
	utils.Loge(err)
}

// ~~ Log Out ~~
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	LogoutUser(w)
	utils.Loge(tplPages.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "Login",
			MainContentURL: "login",
		}))
}
