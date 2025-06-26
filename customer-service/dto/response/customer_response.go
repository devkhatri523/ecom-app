package response

type CustomerResponse struct {
	CustomerID string          `json:"customer_id"`
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Email      string          `json:"email"`
	Address    CustomerAddress `json:"address"`
}
type CustomerAddress struct {
	Street      string `json:"street"`
	HouseNumber string `json:"houseNumber"`
	ZipCode     string `json:"zipCode"`
}
