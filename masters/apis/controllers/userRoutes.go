package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/usecases"
	"testAPI/utils"
)

type UsersHandler struct {
	userUsecase usecases.UserUsecases
}

func UsersController(userUsecase usecases.UserUsecases) *UsersHandler  {
	return &UsersHandler{userUsecase: userUsecase}
}

func (u *UsersHandler) UserAPI(r *mux.Router)  {
	r.HandleFunc("/users", u.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/add", u.InsertUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", u.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", u.UpdateUser).Methods(http.MethodPut)
}

func (u *UsersHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	Users, err := u.userUsecase.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, Users)
	}
}

func (u *UsersHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userID := utils.DecodePathVariable("id", r)
	vars, _:= strconv.Atoi(userID)
	product, err := u.userUsecase.GetUserById(vars)
	if err != nil {
		utils.HandleResponseError(w, http.StatusBadRequest, utils.BAD_REQUEST)
	} else {
		utils.HandleResponse(w, http.StatusOK, product)
	}
}

func (u *UsersHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := utils.JsonDecoder(&user, r)
	fmt.Sprintf("User",user)
	if err != nil {
		utils.HandleRequest(w, http.StatusBadRequest)
	} else {
		err = u.userUsecase.InsertUser(&user)
		if err != nil {
			log.Print(err)
			utils.HandleRequest(w, http.StatusBadGateway)
		} else {
			utils.HandleResponse(w, http.StatusOK, user)
		}
	}
}

func (u *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//var user *models.Users
	//id := utils.DecodePathVariable("sid", r)
	//vars, _:= strconv.Atoi(id)
	//err := utils.JsonDecoder(&user, r)
	//if err != nil {
	//	utils.HandleRequest(w, http.StatusBadRequest)
	//} else {
	//	err = u.userUsecase.UpdateUser(vars, user)
	//	if err != nil {
	//		log.Print(err)
	//		utils.HandleRequest(w, http.StatusBadGateway)
	//	} else {
	//		utils.HandleResponse(w, http.StatusOK, user)
	//	}
	//}
	var user *models.Users
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&user)
	varss, _:= strconv.Atoi(vars["id"])
	err = u.userUsecase.UpdateUser(varss, user)
	if err != nil {
		w.Write([]byte("Data not found !!"))
		log.Println(err)
		return
	}
	byteOfJob, _ := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Update", user))
	// w.Write([]byte("Update successful"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfJob)
}
