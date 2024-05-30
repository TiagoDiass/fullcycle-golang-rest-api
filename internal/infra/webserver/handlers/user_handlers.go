package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/dto"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
)

type UserHandler struct {
	UserDB database.IUserDB
}

func NewUserHandler(userDB database.IUserDB) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
	var createUserDTO dto.CreateUserInput
	err := json.NewDecoder(req.Body).Decode(&createUserDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(createUserDTO.Name, createUserDTO.Email, createUserDTO.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}
