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

type CategoryHandler struct {
	log             *logrus.Logger
	validate        *validator.Validate
	categoryUsecase _interface.CategoryUsecase
}

func NewCategoryHandler(log *logrus.Logger, validate *validator.Validate, categoryUsecase _interface.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{log: log, validate: validate, categoryUsecase: categoryUsecase}
}

func (c *CategoryHandler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryRequest := &request.CategoryCreateRequest{}
	err := helper.ReadFromRequestBody(r, categoryRequest)
	if err != nil {
		c.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	categoryRequest.UserID = userID

	if err = c.validate.Struct(categoryRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	categoryResponse, err := c.categoryUsecase.Create(r.Context(), categoryRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}
	status := http.StatusCreated
	helper.WriteJSON(c.log, w, status, response.WebResponse[*response.CategoryResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   categoryResponse,
	})
}

func (c *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	categoryResponses, err := c.categoryUsecase.GetAll(r.Context(), userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(c.log, w, status, response.WebResponse[[]*response.CategoryResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   categoryResponses,
	})
}

func (c *CategoryHandler) GetDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	categoryIDstr := params.ByName("categoryID")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid category ID",
		})
		return
	}

	categoryResponse, err := c.categoryUsecase.GetByID(r.Context(), categoryID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(c.log, w, status, response.WebResponse[*response.CategoryResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   categoryResponse,
	})
}

func (c *CategoryHandler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryRequest := &request.CategoryUpdateRequest{}
	err := helper.ReadFromRequestBody(r, categoryRequest)
	if err != nil {
		c.log.Errorf("failed to decode request body: %v", err)
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	categoryIDstr := params.ByName("categoryID")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid category ID",
		})
		return
	}

	categoryRequest.UserID = userID
	categoryRequest.ID = categoryID

	if err = c.validate.Struct(categoryRequest); err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	categoryResponse, err := c.categoryUsecase.Update(r.Context(), categoryRequest)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(c.log, w, status, response.WebResponse[*response.CategoryResponse]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   categoryResponse,
	})
}

func (c *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(r.Header.Get("X-User-ID"))
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid user ID",
		})
		return
	}

	categoryIDstr := params.ByName("categoryID")
	categoryID, err := strconv.Atoi(categoryIDstr)
	if err != nil {
		status := http.StatusBadRequest
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  "Invalid category ID",
		})
		return
	}

	err = c.categoryUsecase.Delete(r.Context(), categoryID, userID)
	if err != nil {
		status := helper.DomainErrorToHTTPStatus(err)
		helper.WriteJSON(c.log, w, status, response.WebResponseError{
			Code:   status,
			Status: http.StatusText(status),
			Error:  err.Error(),
		})
		return
	}

	status := http.StatusOK
	helper.WriteJSON(c.log, w, status, response.WebResponse[any]{
		Code:   status,
		Status: http.StatusText(status),
		Data:   nil,
	})
}
