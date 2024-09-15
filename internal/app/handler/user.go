package handler

import (
	"encoding/json"
	"net/http"

	userschema "github.com/K-Kizuku/eisa-auth/internal/app/handler/schema/users"
	"github.com/K-Kizuku/eisa-auth/internal/app/service"
	"github.com/K-Kizuku/eisa-auth/internal/domain/entity"
	"github.com/K-Kizuku/eisa-auth/pkg/errors"
)

type IUserHandler interface {
	SignUp() func(http.ResponseWriter, *http.Request) error
	SignIn() func(http.ResponseWriter, *http.Request) error
}

type UserHandler struct {
	us service.IUserService
}

func NewUserHandler(us service.IUserService) IUserHandler {
	return &UserHandler{us: us}
}

func (h *UserHandler) SignUp() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req userschema.SignUpRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return err
		}
		u := entity.User{
			Username: req.Name,
			Password: req.Password,
			Email:    req.Email,
		}
		createdUser, err := h.us.Create(r.Context(), u)
		if err != nil {
			return err
		}
		token, err := h.us.GenerateJWT(r.Context(), createdUser.ID)
		if err != nil {
			return err
		}
		url, err := h.us.GenerateSignedURL(r.Context(), createdUser.ID)
		if err != nil {
			return err
		}
		res := userschema.SignUpResponse{
			Token:    token,
			SigndURL: url,
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return err
		}
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

func (h *UserHandler) SignIn() func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req userschema.SignInRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return errors.New(http.StatusBadRequest, err)
		}
		id, err := h.us.VerifyPassword(r.Context(), req.Email, req.Password)
		if err != nil {
			return err
		}
		token, err := h.us.GenerateJWT(r.Context(), id)
		if err != nil {
			return err
		}

		res := userschema.SignInResponse{
			Token: token,
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			return errors.New(http.StatusInternalServerError, err)
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
