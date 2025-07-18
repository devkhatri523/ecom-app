package request

type CreateCustomerRequest struct {
	FirstName string          `json:"first_name" validate:"required"`
	LastName  string          `json:"last_name" validate:"required"`
	Email     string          `json:"email" validate:"required"`
	Address   CustomerAddress `json:"address" validate:"required"`
}
type CustomerAddress struct {
	Street      string `json:"street" validate:"required"`
	HouseNumber string `json:"house_number" validate:"required"`
	ZipCode     string `json:"zip_code" validate:"required"`
}

type CustomerUpdateRequest struct {
	Id        int64           `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Email     string          `json:"email"`
	Address   CustomerAddress `json:"address"`
}

type CustomerUpdateAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	ZipCode     string `json:"zip_code"`
}
