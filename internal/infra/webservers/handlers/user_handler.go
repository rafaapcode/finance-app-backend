package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rafaapcode/finance-app-backend/internal/dto"
	"github.com/rafaapcode/finance-app-backend/internal/entity"
	"github.com/rafaapcode/finance-app-backend/internal/infra/database"
)

type UserHandler struct {
	Userdb       database.UserInterface
	JwtExpiresIn int
}

func NewUserHandler(user database.UserInterface) *UserHandler {
	return &UserHandler{
		Userdb: user,
	}
}

func (u *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJwtInput
	jwt := r.Context().Value("token").(*jwtauth.JWTAuth)
	expToken := r.Context().Value("jwtexp").(int)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userData, status, err := u.Userdb.GetUserByEmail(user.Email)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	_, token, err := jwt.Encode(map[string]interface{}{
		"sub": userData.Id.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expToken)).Unix(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser entity.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(newUser.Nome, newUser.Email, newUser.PhotoUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := struct {
			Message string
		}{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	_, status, err := u.Userdb.GetUserByEmail(user.Email)

	if err != nil {
		w.WriteHeader(status)
		errorMessage := struct {
			Message string
		}{Message: err.Error()}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if status == 200 {
		w.WriteHeader(status)
		errorMessage := struct {
			Message string
		}{Message: "Usuáro já existe"}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	status, err = u.Userdb.CreateUser(user)
	if err != nil {
		w.WriteHeader(status)
		return
	}

	w.WriteHeader(status)
}
