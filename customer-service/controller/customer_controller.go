package controller

import (
	"customer-service/dto/request"
	"customer-service/dto/response"
	"customer-service/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	req := request.CreateCustomerRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	customer, err := c.customerService.CreateCustomer(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return

	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   customer,
	}
	ctx.JSON(http.StatusCreated, res)

}

func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	customerId := ctx.Param("customerId")
	custId, _ := strconv.Atoi(customerId)
	_, err := c.customerService.FindByID(int64(custId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	_, err = c.customerService.DeleteCustomer(int64(custId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	respon := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   fmt.Sprintf("Customer with id :%s deleted succssfully", customerId),
	}
	ctx.JSON(http.StatusOK, respon)

}
func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	req := request.CustomerUpdateRequest{}
	err := ctx.ShouldBindJSON(&req)
	id := ctx.Param("customerId")
	customerId, _ := strconv.Atoi(id)
	_, err = c.customerService.UpdateCustomer(int64(customerId), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	respon := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   fmt.Sprintf("Customer with id :%s updated succssfully", customerId),
	}
	ctx.JSON(http.StatusOK, respon)
}
func (c *CustomerController) GetAllCustomers(ctx *gin.Context) {
	data, err := c.customerService.FindAllCustomer()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return

	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *CustomerController) GetCustomerById(ctx *gin.Context) {
	customerId := ctx.Param("customerId")
	custId, _ := strconv.Atoi(customerId)
	data, err := c.customerService.FindByID(int64(custId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}
func (c *CustomerController) IsCustomerExists(ctx *gin.Context) {
	customerId := ctx.Param("customerId")
	custId, _ := strconv.Atoi(customerId)
	isExists, err := c.customerService.IsCustomerExists(int64(custId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	res := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   isExists,
	}
	ctx.JSON(http.StatusOK, res)
}
