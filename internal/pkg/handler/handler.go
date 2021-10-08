/*
 * @Author: Adrian Faisal
 * @Date: 01/10/21 8.19 PM
 */

package handler

import (
	"encoding/json"
	"github.com/apldex/doof/internal/pkg/usecase"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
	HandleGetUserByID(w http.ResponseWriter, r *http.Request)
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
