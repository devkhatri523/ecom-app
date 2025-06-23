package repo

import (
	"customer-service/domain"
	"customer-service/query"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

func NewCustomerRepositoryImpl(Db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{
		Db: Db,
	}
}

func (c CustomerRepositoryImpl) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	result := c.Db.Create(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return customer, nil
}

func (c CustomerRepositoryImpl) FindAllCustomer() ([]domain.Customer, error) {
	var customers []domain.Customer
	stmt := c.Db.Raw(query.FIND_ALL_CUSTOMER_QUERY)
	results := stmt.Find(&customers)
	if results.Error != nil {
		log.Fatal(results.Error)
		return []domain.Customer{}, results.Error
	}
	return customers, nil
}

func (c CustomerRepositoryImpl) UpdateCustomer(customerId int64, customer domain.Customer) (domain.Customer, error) {
	result := c.Db.Where(&domain.Customer{Id: customerId}).UpdateColumns(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
		return domain.Customer{}, errors.New("customer not found")
	}
	return customer, nil

}

func (c CustomerRepositoryImpl) IsCustomerExists(customerId int64) (bool, error) {
	var customer domain.Customer
	whereQuery := fmt.Sprintf("%s where id=%d", query.FIND_ALL_CUSTOMER_QUERY, customerId)
	stmt := c.Db.Raw(whereQuery)
	results := stmt.First(&customer)
	if results.Error != nil {
		return false, results.Error
	}
	if results.RowsAffected == 0 {
		return false, errors.New("customer not found")
	}
	return true, nil

}

func (c CustomerRepositoryImpl) FindByID(customerId int64) (domain.Customer, error) {
	var customer domain.Customer
	whereQuery := fmt.Sprintf("%s where id=%d", query.FIND_ALL_CUSTOMER_QUERY, customerId)
	stmt := c.Db.Raw(whereQuery)
	results := stmt.Find(&customer)
	if results.Error != nil {
		return domain.Customer{}, results.Error
	}
	if results.RowsAffected == 0 {
		return domain.Customer{}, errors.New("customer not found")
	}
	return customer, nil

}

func (c CustomerRepositoryImpl) DeleteCustomer(customerId int64) (int64, error) {

	var customer domain.Customer
	result := c.Db.Where("id=?", customerId).Delete(&customer)
	if result.Error != nil {
		log.Fatal(result.Error)
		return 0, result.Error
	}
	return customerId, nil
}
