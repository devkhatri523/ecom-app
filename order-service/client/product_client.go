package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"order-service/dto/request"
	"order-service/dto/response"

	"github.com/devkhatri523/ecom-app/go-config/v4/config"
	"github.com/devkhatri523/ecom-app/go-service/v2/http"
)

var productServiceUrl = config.Default().GetString("product.service.url")

func PurchaseProducts(orderRequest request.OrderRequest) ([]response.ProductPurchaseResponse, error) {
	jsonRequest, err := json.Marshal(orderRequest.PurchaseProducts)
	purchaseServiceUrl := fmt.Sprintf("%s%s", productServiceUrl, "/purchase") // http;//localhost:8090/api/products/purchase
	if err != nil {
		log.Fatal("Error while parsing payment request", err)
		return []response.ProductPurchaseResponse{}, err
	}
	body := bytes.NewReader(jsonRequest)
	req, err := http.GetHttpPostRequest(purchaseServiceUrl, body)
	if err != nil {
		log.Fatal("Error while creating payment post request", err)
		return []response.ProductPurchaseResponse{}, err
	}
	res, err := http.ExecuteHttpRequest(req)
	if err != nil {
		log.Fatal("Error while executing product purchase post request", err)
		return []response.ProductPurchaseResponse{}, err
	}
	var purchaseProductsResponse response.PurchaseData
	err = http.ReadHttpBodyAsJson(&purchaseProductsResponse, res)
	if err != nil {
		log.Fatal("Error while parsing product purchase response ", err)
		return []response.ProductPurchaseResponse{}, err
	}
	return purchaseProductsResponse.Data, nil

}
