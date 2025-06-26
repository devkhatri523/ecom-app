package service

import (
	"customer-service/domain"
	"customer-service/dto/request"
	"customer-service/dto/response"
	"customer-service/repo"
	"strconv"
)

type CustomerServiceImpl struct {
	repository repo.CustomerRepository
}

func NewCustomerServiceImpl(repository repo.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		repository: repository,
	}
}

func (c CustomerServiceImpl) CreateCustomer(customerRequest request.CreateCustomerRequest) (response.CustomerResponse, error) {
	customer := domain.Customer{
		FirstName: customerRequest.FirstName,
		LastName:  customerRequest.LastName,
		Email:     customerRequest.Email,
		Address: domain.Address{
			Street:      customerRequest.Address.Street,
			HouseNumber: customerRequest.Address.HouseNumber,
			ZipCode:     customerRequest.Address.ZipCode,
		},
	}
	res, err := c.repository.CreateCustomer(customer)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	customerResponse := response.CustomerResponse{
		CustomerID: strconv.FormatInt(res.Id, 10),
		FirstName:  res.FirstName,
		LastName:   res.LastName,
		Email:      res.Email,
		Address: response.CustomerAddress{
			Street:      res.Address.Street,
			HouseNumber: res.Address.HouseNumber,
			ZipCode:     res.Address.ZipCode,
		},
	}
	return customerResponse, nil

}

func (c CustomerServiceImpl) FindAllCustomer() ([]response.CustomerResponse, error) {
	var cutomerRes []response.CustomerResponse
	res, err := c.repository.FindAllCustomer()
	if err != nil {
		return nil, err
	}
	for _, value := range res {
		customer := response.CustomerResponse{
			CustomerID: strconv.FormatInt(value.Id, 10),
			FirstName:  value.FirstName,
			LastName:   value.LastName,
			Email:      value.Email,
			Address: response.CustomerAddress{
				Street:      value.Address.Street,
				HouseNumber: value.Address.HouseNumber,
				ZipCode:     value.Address.ZipCode,
			},
		}
		cutomerRes = append(cutomerRes, customer)
	}
	return cutomerRes, nil

}

func (c CustomerServiceImpl) UpdateCustomer(customerId int64, cutomerUpdateRequest request.CustomerUpdateRequest) (string, error) {
	customerUpdateModel := domain.Customer{
		FirstName: cutomerUpdateRequest.FirstName,
		LastName:  cutomerUpdateRequest.LastName,
		Email:     cutomerUpdateRequest.Email,
		Address: domain.Address{
			Street:      cutomerUpdateRequest.Address.Street,
			HouseNumber: cutomerUpdateRequest.Address.HouseNumber,
			ZipCode:     cutomerUpdateRequest.Address.ZipCode,
		},
	}
	_, err := c.repository.UpdateCustomer(customerId, customerUpdateModel)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(customerId, 10), nil
}

func (c CustomerServiceImpl) IsCustomerExists(customerId int64) (bool, error) {
	isAlreadyExists, err := c.repository.IsCustomerExists(customerId)
	if err != nil {
		return false, err
	}
	return isAlreadyExists, nil
}

func (c CustomerServiceImpl) FindByID(customerId int64) (response.CustomerResponse, error) {
	result, err := c.repository.FindByID(customerId)
	if err != nil {
		return response.CustomerResponse{}, err
	}
	res := response.CustomerResponse{
		CustomerID: strconv.FormatInt(result.Id, 10),
		FirstName:  result.FirstName,
		LastName:   result.LastName,
		Email:      result.Email,
		Address: response.CustomerAddress{
			Street:      result.Address.Street,
			HouseNumber: result.Address.HouseNumber,
			ZipCode:     result.Address.ZipCode,
		},
	}
	return res, nil

}

func (c CustomerServiceImpl) DeleteCustomer(customerId int64) (int64, error) {
	customerId, err := c.repository.DeleteCustomer(customerId)
	if err != nil {
		return 0, err
	}
	return customerId, nil
}
