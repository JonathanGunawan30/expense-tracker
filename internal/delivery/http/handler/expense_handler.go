package handler

import (
	"jonathangunawan30/expense-tracker/internal/delivery/http/helper"
	"jonathangunawan30/expense-tracker/internal/domain/entity/request"
	"jonathangunawan30/expense-tracker/internal/domain/entity/response"
	_interface "jonathangunawan30/expense-tracker/internal/usecase/interface"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type ExpenseHandler struct {
	log            *logrus.Logger
	validate       *validator.Validate
	expenseUsecase _interface.ExpenseUsecase
}

func NewExpenseHandler(log *logrus.Logger, validate *validator.Validate, expenseUsecase _interface.ExpenseUsecase) *ExpenseHandler {
	return &ExpenseHandler{log: log, validate: validate, expenseUsecase: expenseUsecase}
}

func (e *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	expenseRequest := &request.ExpenseCreateRequest{}
	err := helper.ReadFromRequestBody(r, expenseRequest)
	if err != nil {
		e.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	expenseRequest.UserID = userID

	if err = e.validate.Struct(expenseRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	expenseResponse, err := e.expenseUsecase.Create(r.Context(), expenseRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusCreated
	helper.WriteJSON(e.log, w, status, response.WebResponse[*response.ExpenseResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   expenseResponse,
	})
}

func (e *ExpenseHandler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	expenseResponses, err := e.expenseUsecase.GetAll(r.Context(), userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(e.log, w, status, response.WebResponse[[]*response.ExpenseResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   expenseResponses,
	})
}

func (e *ExpenseHandler) GetDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	expenseIDstr := params.ByName("expenseID")
	expenseID, err := strconv.Atoi(expenseIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid expense ID",
		})
		return
	}

	expenseResponse, err := e.expenseUsecase.GetDetail(r.Context(), expenseID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(e.log, w, status, response.WebResponse[*response.ExpenseResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   expenseResponse,
	})
}

func (e *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	expenseRequest := &request.ExpenseUpdateRequest{}
	err := helper.ReadFromRequestBody(r, expenseRequest)
	if err != nil {
		e.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	expenseIDstr := params.ByName("expenseID")
	expenseID, err := strconv.Atoi(expenseIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid expense ID",
		})
		return
	}

	expenseRequest.UserID = userID
	expenseRequest.ID = expenseID

	if err = e.validate.Struct(expenseRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	expenseResponse, err := e.expenseUsecase.Update(r.Context(), expenseRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(e.log, w, status, response.WebResponse[*response.ExpenseResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   expenseResponse,
	})
}

func (e *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	expenseIDstr := params.ByName("expenseID")
	expenseID, err := strconv.Atoi(expenseIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid expense ID",
		})
		return
	}

	err = e.expenseUsecase.Delete(r.Context(), expenseID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(e.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(e.log, w, status, response.WebResponse[any]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   nil,
	})
}
