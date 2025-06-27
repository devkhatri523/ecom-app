package client

import (
	"fmt"
	"log"
	"payment-service/dto/response"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-service/v2/http"
)

var customerServiceUrl = config.Default().GetString("customer.service.url")

func GetCustomerDetails(customerId string) (response.CustomerResponse, error) {
	endPoint := fmt.Sprintf("%s/%s", customerServiceUrl, customerId) /// http://localhost:8080/api/customer/123
	req, err := http.GetHttpGetRequest(endPoint, nil)
	if err != nil {
		log.Fatal("error while creating http get request ", err)
		return response.CustomerResponse{}, err
	}
	res, err := http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal("error while  executing  cutomer  request", err)
		return response.CustomerResponse{}, err
	}
	var customer response.CustomerData
	err = http.ReadHttpBodyAsJson(&customer, res)
	if err != nil {
		log.Fatal("error while  parsing  cutomer  response", err)
		return response.CustomerResponse{}, err

	}
	return customer.Data, nil

}
