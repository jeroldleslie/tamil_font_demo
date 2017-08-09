package controllers

import (
	"tamil_font_demo/commons"
	"tamil_font_demo/models"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"io/ioutil"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"os/user"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

//const (
// For simplicity these files are in the same folder as the app binary.
// You shouldn't do this in production.
//privKeyPath = "~/.tamil_font_demo/keys/app.rsa"
//pubKeyPath  = "~/.tamil_font_demo/keys/app.rsa.pub"
//)

var (
	privKeyPath = "app.rsa"
	pubKeyPath  = "app.rsa.pub"
	verifyKey   *rsa.PublicKey
	signKey     *rsa.PrivateKey
)

var VerifyKey, SignKey []byte

type AuthController struct {
	session  *mgo.Session
	response *commons.ResponseController
}

func NewAuthController(s *mgo.Session) *AuthController {
	ac := &AuthController{
		session:  s,
		response: commons.NewResponseController(),
	}

	initKeys()
	//ensureIndex(u.session)
	return ac
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initKeys() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := usr.HomeDir
	fmt.Print(homeDir)
	privKeyPath = "src/tamil_font_demo/resources/keys/" + privKeyPath
	pubKeyPath = "src/tamil_font_demo/resources/keys/" + pubKeyPath
	signBytes, err := ioutil.ReadFile(privKeyPath)
	fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)
}

func (ac AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user models.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	s := ac.session.Copy()
	defer s.Close()
	c := s.DB("tamil_font_demo").C("users")
	var dbuser models.User
	err = c.Find(bson.M{"username": user.Username}).One(&dbuser)

	if err != nil {
		ac.response.WriteError(w, "User not found", http.StatusNotFound)
		log.Println("Failed get user: ", err)
		return
	}

	if user.Password != dbuser.Password {
		fmt.Println("Invalid credentials")
		ac.response.WriteError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(72)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = dbuser.Username
	claims["type"] = dbuser.Type

	token.Claims = claims

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		fatal(err)
	}

	tokenString, err := token.SignedString(signKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		fatal(err)
	}

	response := Token{tokenString}
	ac.response.WriteSuccess(w, response, http.StatusOK)

}

func (ac AuthController) ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}
