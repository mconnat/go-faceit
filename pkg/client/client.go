package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FaceITClient struct {
	HTTPClient  *http.Client
	BearerToken string
	BaseURL     string
}

func New(bearer string) (FaceITClient, error) {
	var client FaceITClient
	if bearer == "" {
		return client, fmt.Errorf("empty Bearer token")
	}
	client.BearerToken = bearer
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS13,
		},
	}
	httpClient := http.Client{Transport: tr}
	client.HTTPClient = &httpClient
	client.BaseURL = "https://open.faceit.com/data/v4"
	return client, nil
}

func Get[T any](responseModel T, c *FaceITClient, endpoint string, params map[string]interface{}) (T, error) {
	url := c.BaseURL + endpoint
	req, err := http.NewRequest("GET", url, nil)

	if params != nil {

		q := req.URL.Query()
		for k, v := range params {
			switch v.(type) {
			case string:
				q.Add(k, fmt.Sprintf("%v", v))
			case []string:
				values := v.([]string)
				for _, p := range values {
					q.Add(k, fmt.Sprintf("%v", p))
				}
			default:
				return responseModel, fmt.Errorf("unknow type for params, should be string or []string")
			}

		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Add("Authorization", "Bearer "+c.BearerToken)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return responseModel, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseModel, err
	}
	defer resp.Body.Close()
	jsonByte := body
	err = json.Unmarshal(jsonByte, &responseModel)
	if err != nil {
		return responseModel, err
	}
	return responseModel, nil
}
