package noundo

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/schema"
)

var jwtCookieKey string

var hmacSecret []byte

var decoder = schema.NewDecoder()

func init() {
	stringSecret := os.Getenv("FORUM_JWT_HMAC_SECRET")
	if len(stringSecret) == 0 {
		stringSecret = "hmacSampleSecret"
	}
	hmacSecret = []byte(stringSecret)

	// Its for a `Key: Value` pair. Just a random key so no collisions will occur
	jwtCookieKey = "bueykivxivcxf436ugkhgu8owy3886^$&7ae"
}

// Decode `http.Request.Body` into `AuthMe struct`
func GetLoginMe(r *http.Request) (*LoginMe, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, err
	}
	var loginMe LoginMe
	err = decoder.Decode(&loginMe, r.Form)
	return &loginMe, err
}

// Function for checking if
// Read AuthMe doc
func ValidateAuthMe(a *LoginMe) (bool, error) {
	return true, nil
}

// Parse http.Request.Body, then check if data passed is valid and if so set
// header on ResponseWriter to JWT with appropriate data.
func LoginUser(auth Authenticator, w http.ResponseWriter, r *http.Request) error {
	authMe, err := GetLoginMe(r)
	if err != nil {
		log.Printf("parse Credentials from request: %s\n", err)
		return err
	}

	err = auth.ValidateAuthMe(authMe)
	if err != nil {
		log.Printf("validate Credentials: %s\n", err)
		return err
	}

	user := auth.GetUserByEmail(authMe.Email)

	// TODO: currently unsecure
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		newJWTMapClaims(JWTFields{Username: user.Username()}),
	)

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Printf("signing token error: %s\n", err)
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     jwtCookieKey,
		Value:    tokenString,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode, // TODO: Unsecure?
	})

	return nil
}

func LogoutUser(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     jwtCookieKey,
		Value:    "",
		HttpOnly: true,
	})
}

func GetRegisterMe(r *http.Request) (*RegisterMe, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	var registerMe RegisterMe
	err = decoder.Decode(&registerMe, r.Form)
	return &registerMe, err
}

func RegisterUser(auth Authenticator, r *http.Request) *RegisterMeResponse {
	registerMe, err := GetRegisterMe(r)

	if err != nil {
		return &RegisterMeResponse{RestResp{false, DecodeErr}}
	}

	return auth.RegisterUser(registerMe)
}

// Validate the JWT sent with the incoming Request
func JWTCheckAndParse(r *http.Request) (JWTFields, error) {
	var jfEmpty JWTFields
	c, err := r.Cookie(jwtCookieKey)
	if err != nil {
		return jfEmpty, err
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	// token.Valid is (I think) redudant, because it is set just before `return` from `jwt.ParseWithClaims` - which is called by `jwt.Parse`.
	if err != nil || !token.Valid {
		return jfEmpty, err
	}

	// Type assertion. Can be performed on interfaces only.
	// token.Claims is an interface
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jfEmpty, errors.New("failed type assertion: `token.Claims.(jwt.MapClaims)`")
	}

	jf, err := jwtFieldsFromMapClaims(mapClaims)
	if err != nil {
		return jfEmpty, err
	}

	return jf, nil
}
