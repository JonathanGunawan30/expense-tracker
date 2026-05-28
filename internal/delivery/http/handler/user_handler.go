package handler

import (
	"jonathangunawan30/expense-tracker/internal/delivery/http/helper"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	_interface "jonathangunawan30/expense-tracker/internal/usecase/interface"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	log         *logrus.Logger
	validate    *validator.Validate
	userUsecase _interface.UserUsecase
}

func NewUserHandler(log *logrus.Logger, validate *validator.Validate, userUsecase _interface.UserUsecase) *UserHandler {
	return &UserHandler{log: log, validate: validate, userUsecase: userUsecase}
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userRequest := &request.UserRegisterRequest{}
	err := helper.ReadFromRequestBody(r, userRequest)
	if err != nil {
		u.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(u.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	if err = u.validate.Struct(userRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(u.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userResponse, err := u.userUsecase.Register(r.Context(), userRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(u.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}
	status := http.StatusCreated
	helper.WriteJSON(u.log, w, status, response.WebResponse[*response.UserResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   userResponse,
	})
}

func (u *UserHandler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userResponses, err := u.userUsecase.GetAll(r.Context())
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(u.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(u.log, w, status, response.WebResponse[[]*response.UserResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   userResponses,
	})
}
