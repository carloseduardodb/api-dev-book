package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("register"); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("name_or_nick"))
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	users, err := repository.Find(nameOrNick, tokenUserID)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	user, err := repository.FindById(id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if id != tokenUserID {
		responses.ERROR(w, http.StatusForbidden, errors.New("you can only update your own account"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare("update"); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user.ID = id
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Update(user); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if id != tokenUserID {
		responses.ERROR(w, http.StatusForbidden, errors.New("you can only delete your own account"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Delete(id); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if id == tokenUserID {
		responses.ERROR(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Follow(id, tokenUserID); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if id == tokenUserID {
		responses.ERROR(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	if err = repository.Unfollow(id, tokenUserID); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	followers, err := repository.FindFollowers(id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	following, err := repository.FindFollowing(id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, following)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var pass models.ChangePassword
	if err := json.Unmarshal(body, &pass); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewRepositoryUsers(db)
	userDatabase, err := repository.FindByID(tokenUserID)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}
	if err := security.CheckPasswordHash(userDatabase.Password, pass.Password); err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("password is incorrect"))
		return
	}
	hashedNewPassword, err := security.Hash(pass.NewPassword)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	userDatabase.Password = string(hashedNewPassword)
	if err := repository.UpdatePassword(userDatabase); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
