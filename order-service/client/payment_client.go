package client

import (
	"bytes"
	"encoding/json"
	"log"
	"order-service/dto/request"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-service/v2/http"
)

var paymentServiceUrl = config.Default().GetString("payment.service.url")

func RequestPayment(paymentRequest request.PaymentRequest) error {
	jsonRequest, err := json.Marshal(paymentRequest)
	if err != nil {
		log.Fatal("Error while parsing payment request", err)
		return err
	}
	body := bytes.NewReader(jsonRequest)
	req, err := http.GetHttpPostRequest(paymentServiceUrl, body)
	if err != nil {
		log.Fatal("Error while creating payment post request", err)
		return err
	}
	_, err = http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal("Error while executing payment post request", err)
		return err
	}
	return nil

}
