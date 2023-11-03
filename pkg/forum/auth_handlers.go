package forum

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
	login "github.com/reddec/go-login"
)

const jwtCookieKey = "bueykivxivcxf436ugkhgu8owy3886^$&7ae"

var hmacSecret string

func init() {
	hmacSecret = "hmacSampleSecret"
}

func loginFunc(writer http.ResponseWriter, r *http.Request, cred login.UserPassword) error {
	ok := cred.User == "admin" && cred.Password == "admin" // user proper login and validation
	if !ok {
		return fmt.Errorf("username or password is incorrect")
	}
	// use sessions/JWT/cookies and mark following requests as authorized
	return nil
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var login LoginRequestPostBody

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: currently unsecure
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newJWTMapClaims(JWTFields{Login: login.Login}))

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     jwtCookieKey,
		Value:    tokenString,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode, // TODO: Unsecure?
	})

	http.Redirect(w, r, "/", 200)
}

func CheckAndParseJWT(w http.ResponseWriter, r *http.Request) (JWTFields, error) {
	var jf JWTFields
	c, err := r.Cookie(jwtCookieKey)
	if err != nil {
		return jf, err
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return jf, err
	}

	// Type assertion. Can be performed on interfaces only.
	// token.Claims is an interface
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	if token.Valid {
	}
	return token, nil
}

func ListenAndServe() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<html><body><h1>Home</h1><br/><a href='/login'>Login</a></body></html>"))
	})

	http.Handle("/login", login.New[login.UserPassword](loginFunc, login.Log(func(err error) {
		log.Println(err)
	})))
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		u1 := User{
			Id:       NewRandId(),
			Username: "kacper",
		}
		p1 := Post{
			Postable: Postable{
				Id:            NewRandId(),
				UserId:        u1.Id,
				Contents:      "wubba lubba dab dab",
				TimeStampable: TimeStampable{Timestamp: 0}},
			Reactionable: Reactionable{Reactions: []Reaction{}},
		}
		RenderPost(w, &p1)
	})
	const addr = "127.0.0.1:8083"
	log.Println("listening on: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
