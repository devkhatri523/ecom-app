package service

import (
	"customer-service/dto/request"
	"customer-service/dto/response"
)

type CustomerService interface {
	CreateCustomer(customerRequest request.CreateCustomerRequest) (response.CustomerResponse, error)
	FindAllCustomer() ([]response.CustomerResponse, error)
	UpdateCustomer(customerId int64, cutomerUpdateRequest request.CustomerUpdateRequest) (string, error)
	IsCustomerExists(customerId int64) (bool, error)
	FindByID(customerId int64) (response.CustomerResponse, error)
	DeleteCustomer(customerId int64) (int64, error)
}
