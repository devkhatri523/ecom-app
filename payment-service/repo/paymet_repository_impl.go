package repo

import (
	"errors"
	"payment-service/domain"

	"gorm.io/gorm"
)

type PaymetRepositoryImpl struct {
	Db *gorm.DB
}

func NewPaymentRepostioryImpl(db *gorm.DB) PaymentRepository {
	return &PaymetRepositoryImpl{Db: db}
}

func (p PaymetRepositoryImpl) CreatePayment(payment domain.Payment) (int32, error) {
	result := p.Db.Create(&payment)
	if result.Error != nil {
		return 0, errors.New(result.Error.Error())
	}
	return payment.Id, nil
}
