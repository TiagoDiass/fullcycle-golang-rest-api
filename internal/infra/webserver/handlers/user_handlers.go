package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/dto"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/entity"
	"github.com/TiagoDiass/fullcycle-golang-rest-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.IUserDB
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.IUserDB, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

// CreateUser godoc
// @Summary      Creates an user
// @Description  Creates an user with name, email and password.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request	body			dto.CreateUserInput		true		"user request"
// @Success      200  		{object}  entity.User
// @Failure      500  		{object}  Error
// @Router       /users 	[post]
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
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}

// CreateSession godoc
// @Summary      Creates a session
// @Description  Creates a session using an user email and password.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request	body			dto.CreateSessionInput		true		"user credentials"
// @Success      200  		{object}  dto.CreateSessionOutput
// @Failure      400  		{object}  Error
// @Failure      401  		{object}  Error
// @Router       /session [post]
func (h *UserHandler) CreateSession(w http.ResponseWriter, req *http.Request) {
	var createSessionDTO dto.CreateSessionInput
	err := json.NewDecoder(req.Body).Decode(&createSessionDTO)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(createSessionDTO.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	passwordsMatch := user.ValidatePassword(createSessionDTO.Password)

	if !passwordsMatch {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "Unauthorized"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.CreateSessionOutput{AccessToken: tokenString}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	return
}
