package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
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

func Get[T any](responseModel T, c *FaceITClient, endpoint string, params any) (T, error) {
	url := c.BaseURL + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return responseModel, err
	}

	if params != nil {
		query, err := buildQueryString(req, params)
		if err != nil {
			return responseModel, err
		}
		req.URL.RawQuery = query
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

func buildQueryString(req *http.Request, v interface{}) (string, error) {
	isNil, err := isZero(v)
	if err != nil {
		return "", err
	}

	if isNil {
		return "", nil
	}

	query := req.URL.Query()
	vType := reflect.TypeOf(v).Elem()
	vValue := reflect.ValueOf(v).Elem()

	for i := 0; i < vType.NumField(); i++ {
		var defaultValue string

		field := vType.Field(i)
		tag := field.Tag.Get("query")

		if tag == "" {
			continue
		}

		// Get the default value from the struct tag
		if strings.Contains(tag, ",") {
			tagSlice := strings.Split(tag, ",")

			tag = tagSlice[0]
			defaultValue = tagSlice[1]
		}

		if field.Type.Kind() == reflect.Slice {
			// Attach any slices as query params
			fieldVal := vValue.Field(i)
			for j := 0; j < fieldVal.Len(); j++ {
				query.Add(tag, fmt.Sprintf("%v", fieldVal.Index(j)))
			}
		} else {
			// Add any scalar values as query params
			fieldVal := fmt.Sprintf("%v", vValue.Field(i))

			// If no value was set by the user, use the default
			// value specified in the struct tag.
			if fieldVal == "" || fieldVal == "0" {
				if defaultValue == "" {
					continue
				}

				fieldVal = defaultValue
			}

			query.Add(tag, fieldVal)
		}
	}

	return query.Encode(), nil
}

func isZero(v interface{}) (bool, error) {
	t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil
}
