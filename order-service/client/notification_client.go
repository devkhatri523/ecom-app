package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"order-service/dto/request"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-service/v2/http"
)

var notificationServiecUrl = config.Default().GetString("notification.service.url")

func SendOrderConfirmation(request request.OrderConfirmation) error {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		log.Fatal("Error while parsing order confirmation request", err)
		return err
	}
	var notificationUrl = fmt.Sprintf("%s%s", notificationServiecUrl, "/orderemail")
	body := bytes.NewReader(jsonRequest)
	req, err := http.GetHttpPostRequest(notificationUrl, body)
	if err != nil {
		log.Fatal("Error while creating notification post request", err)
		return err
	}
	_, err = http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal("Error while executing notification post request", err)
		return err
	}
	return nil
}
