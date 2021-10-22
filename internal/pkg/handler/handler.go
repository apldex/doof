/*
 * @Author: Adrian Faisal
 * @Date: 01/10/21 8.19 PM
 */

package handler

import (
	"encoding/json"
	"github.com/apldex/doof/internal/pkg/usecase"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
	HandleGetUserByID(w http.ResponseWriter, r *http.Request)
	HandleCreateUserForm(w http.ResponseWriter, r *http.Request)
	HandleCreateUser(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	uc usecase.Usecase
}

// New creates a new handler
func New(uc usecase.Usecase) Handler {
	return &handler{uc: uc}
}

func (h *handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (h *handler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("[handler.handleGetUserByID] convert string to int failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	user, err := h.uc.GetUserByID(userID)
	if err != nil {
		log.Printf("[handler.handleGetUserByID] unable to find user by id: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	b, err := json.Marshal(&user)
	if err != nil {
		log.Printf("[handler.handleGetUserByID] error while marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("internal server error"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func (h *handler) HandleCreateUserForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("[HandleCreateUser] parse form failed: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad request"))

		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	err = h.uc.CreateUser(name, email)
	if err != nil {
		log.Printf("[HandleCreateUser] error create user: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("error create user"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))
}

func (h *handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	requestData, err := h.readCreateUserRequest(r)
	if err != nil {
		log.Printf("[HandleCreateUser] error read request data: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("error read request data"))

		return
	}

	err = h.uc.CreateUser(requestData.Name, requestData.Email)
	if err != nil {
		log.Printf("[HandleCreateUser] error create user: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("error create user"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"message": "ok"}`))
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *handler) readCreateUserRequest(r *http.Request) (*createUserRequest, error) {
	req := createUserRequest{}

	switch r.Header.Get("Content-Type") {
	case "application/json":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &req)
		if err != nil {
			return nil, err
		}
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			return nil, err
		}

		req.Name = r.FormValue("name")
		req.Email = r.FormValue("email")
	}

	return &req, nil
}
