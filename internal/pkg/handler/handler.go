/*
 * @Author: Adrian Faisal
 * @Date: 01/10/21 8.19 PM
 */

package handler

import "net/http"

type Handler interface {
	HandleHealthCheck(w http.ResponseWriter, r *http.Request)
}

type handler struct {
}

// New creates a new handler
func New() Handler {
	return &handler{}
}

func (h *handler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
