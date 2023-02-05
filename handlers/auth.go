package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	authdto "ptedi/dto/auth"
	dto "ptedi/dto/result"
	"ptedi/models"
	"ptedi/pkg/bcrypt"
	jwtToken "ptedi/pkg/jwt"
	"ptedi/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.Id_User
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)

	if errGenerateToken != nil {
		fmt.Println(errGenerateToken)
		return
	}

	loginResponse := authdto.LoginResponse{
		Id:       user.Id_User,
		Username: user.Username,
		Fullname: user.Fullname,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	user, err := h.AuthRepository.CekUser(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "User Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userData := authdto.AuthResponse{
		Id:       user.Id_User,
		Username: user.Username,
		Fullname: user.Fullname,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: userData}
	json.NewEncoder(w).Encode(response)
}
