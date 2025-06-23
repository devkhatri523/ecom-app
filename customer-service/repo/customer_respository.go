package repo

import "customer-service/domain"

type CustomerRepository interface {
	CreateCustomer(customer domain.Customer) (domain.Customer, error)
	FindAllCustomer() ([]domain.Customer, error)
	UpdateCustomer(customerId int64, customer domain.Customer) (domain.Customer, error)
	IsCustomerExists(customerId int64) (bool, error)
	FindByID(customerId int64) (domain.Customer, error)
	DeleteCustomer(customerId int64) (int64, error)
}
