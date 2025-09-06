package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api-em-go/internal/dto"
	"api-em-go/internal/models"
	"api-em-go/internal/services"

	"github.com/gorilla/mux"
)

// UserHandler lida com as requisições HTTP relacionadas a usuários.
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler cria uma nova instância de UserHandler.
func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{userService: s}
}

// CreateUser lida com a criação de um novo usuário.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.userService.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// GetUserByID lida com a busca de um usuário por ID.
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID de usuário inválido", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	res := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
