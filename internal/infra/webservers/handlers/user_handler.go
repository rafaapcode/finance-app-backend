package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
	"golang.org/x/oauth2"
)

type googleAuth struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type UserHandler struct {
	Userdb database.UserInterface
	App    *oauth2.Config
}

func NewUserHandler(user database.UserInterface, googleApp *oauth2.Config) *UserHandler {
	return &UserHandler{
		Userdb: user,
		App:    googleApp,
	}
}

func (u *UserHandler) getJwt(email string, JwtExpiresIn int, jwt *jwtauth.JWTAuth) (string, error) {
	userData, _, err := u.Userdb.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	_, token, err := jwt.Encode(map[string]interface{}{
		"sub": userData.Id.String(),
		"exp": time.Now().Add(time.Second * time.Duration(JwtExpiresIn)).Unix(),
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserHandler) create(newUser *entity.User) (int, error) {

	_, status, err := u.Userdb.GetUserByEmail(newUser.Email)

	if status == 200 {
		return status, err
	}

	status, err = u.Userdb.CreateUser(newUser)
	if err != nil {
		return status, err
	}

	return status, nil
}

func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {
	url := u.App.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (u UserHandler) CallbackAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	t, err := u.App.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client := u.App.Client(context.Background(), t)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var userData googleAuth
	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser, err := entity.NewUser(userData.Name, userData.Email, userData.Picture)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = newUser.Validate()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := u.create(newUser)

	if err != nil {
		fmt.Println("Erro aqui")
		http.Error(w, err.Error(), status)
		return
	}
	expJwt := r.Context().Value("jwtexp").(int)
	jwt := r.Context().Value("token").(*jwtauth.JWTAuth)

	token, err := u.getJwt(newUser.Email, expJwt, jwt)
	fmt.Print(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		Access_Token string `json:"access_token"`
	}{Access_Token: token}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(accessToken)
}
