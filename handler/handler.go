package handler

import (
	"encoding/json"
	"go.mod/service"
	"io/ioutil"
	"net/http"
	"strings"
)

//Handler package interface
type IWalletHandler interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	WalletMethods(w http.ResponseWriter, r *http.Request)
}

//Handler package struct implements service package's interface.
type WalletHandler struct {
	service service.IUserService
}

func (h *WalletHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	//using service's function, it returns all users
	users := h.service.GetAllUsers()

	//turn that list into JSON
	jsonBytes, err := json.Marshal(users)
	//if error is not nil, sends 500 HTTP response, and write an error.
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	//added http header values.
	w.Header().Add("content-type", "application/json")
	//sends 200 HTTP response
	w.WriteHeader(http.StatusOK)
	//write the users([]byte)
	w.Write(jsonBytes)

}

func (h *WalletHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	//it returns what comes after / path
	username := strings.TrimPrefix(r.URL.Path, "/")

	//using service's function, it returns wanted user
	user, err := h.service.GetUser(username)
	if err != nil {
		//404 HTTP response
		w.WriteHeader(http.StatusNotFound)
	}

	//turn that user list into JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		//500 HTTP response
		w.WriteHeader(http.StatusInternalServerError)
		//write an error
		w.Write([]byte(err.Error()))
	}
	//added http header values.
	w.Header().Add("content-type", "application/json")
	//sends 200 HTTP response
	w.WriteHeader(http.StatusOK)
	//write the users([]byte)
	w.Write(jsonBytes)
}

func (h *WalletHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	//it returns what comes after / path
	username := strings.TrimPrefix(r.URL.Path, "/")

	//using service's function, it creat a user
	user := h.service.CreateUser(username)

	//turn that user list into JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//sends 201 HTTP response
	w.WriteHeader(http.StatusCreated)
	//write the users([]byte)
	w.Write(jsonBytes)
}

func (h *WalletHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//it returns what comes after / path
	username := strings.TrimPrefix(r.URL.Path, "/")

	//reads request body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	//close body, after the execution.
	defer r.Body.Close()

	if err != nil {
		//sends 500 HTTP response
		w.WriteHeader(http.StatusInternalServerError)
		//write an error
		w.Write([]byte(err.Error()))
		return
	}

	//make a map value
	user := make(map[string]int)
	//turn that json byte into user map
	err = json.Unmarshal(bodyBytes, &user)

	if err != nil {
		//sends 500 HTTP response
		w.WriteHeader(http.StatusInternalServerError)
		//write an error
		w.Write([]byte(err.Error()))
		return
	}

	//initialize balance value
	balance := user["balance"]
	//using service's function, it returns updated user
	h.service.UpdateUser(username, balance)
}

func (h *WalletHandler) WalletMethods(w http.ResponseWriter, r *http.Request) {
	//it returns what comes after / path
	username := strings.TrimPrefix(r.URL.Path, "/")

	//method switch-case
	switch {
	case r.Method == "GET" && username == "":
		h.GetAllUsers(w, r)
		return
	case r.Method == "GET" && username != "":
		h.GetUser(w, r)
	case r.Method == "PUT" && username != "":
		h.CreateUser(w, r)
		return
	case r.Method == "POST" && username != "":
		h.UpdateUser(w, r)
		return
	default:
		//default 405 HTTP
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

//constructure function
func NewWalletHandler(s service.IUserService) IWalletHandler {
	return &WalletHandler{service: s}
}
