package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetHttpGetRequest(url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("GET", url, body)
}

func GetHttpPostRequest(url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("POST", url, body)
}

func GetHttpPatchRequest(url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("PATCH", url, body)
}
func GetHttpPutHRequest(url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("PUT", url, body)
}

func GetHttpDeleteHRequest(url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("DELETE", url, body)
}

func ExecuteHttpRequest(req *http.Request) (*http.Response, error) {
	c := http.Client{}
	defer c.CloseIdleConnections()
	return c.Do(req)
}

func ReadHttpBodyAsJson[a interface{}](obj *a, res *http.Response) error {
	b, err := ReadHttpBodyAsBytes(res)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj)
	if err != nil {
		return err
	}
	return nil

}
func ReadHttpBodyAsString[a interface{}](res *http.Response) (string, error) {
	b, err := ReadHttpBodyAsBytes(res)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
func CreateHeadersMap(key string, value string) map[string]string {
	return map[string]string{
		key: value,
	}
}

func SetHeaders(req *http.Request, headers map[string]string) {

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func GetAuthorizationHeaderMap(authorization string) map[string]string {
	return CreateHeadersMap("Authorization", authorization)
}

func GetJsonContentTypeHeaderMap(authorization string) map[string]string {
	hm := CreateHeadersMap("Content-Type", "application/json")
	hm["Accept"] = "application/json"
	return hm
}

func GetAuthorizationJsonContentTypeHeaderMap(authorization string) map[string]string {
	hm := CreateHeadersMap("Content-Type", "application/json")
	hm["Accept"] = "application/json"
	hm["Authorization"] = authorization
	return hm
}
func ReadHttpBodyAsBytes(res *http.Response) ([]byte, error) {
	if res == nil {
		return nil, errors.New("Response is nil")
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, err

}
