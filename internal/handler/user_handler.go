package handler

import (
	"encoding/json"
	"net/http"

	"pipelines/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {

	users := h.service.GetUsers()

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(users)
}
