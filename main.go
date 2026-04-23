package chippg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/asaskevich/govalidator"
)

const BASE_URL = "https://gate.chip-in.asia/api/v1"

// Client config
type ClientConfig struct {
	BrandID string `json:"brand_id" encore:"required"`
	Secret  string `json:"secret" encore:"required"`
}

func NewClient(request ClientConfig) (ClientConfig, error) {
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return ClientConfig{}, err
	}

	return request, nil
}

func (c *ClientConfig) CreatePurchase(request PurchaseOption) (*PurchaseResponse, error) {
	url := BASE_URL + "/purchases/"

	reqBody := createPurchaseRequest{
		Client:                 request.Client,
		Purchase:               request.Purchase,
		BrandID:                c.BrandID,
		SuccessRedirect:        request.SuccessRedirect,
		FailureRedirect:        request.FailureRedirect,
		CancelRedirect:         request.CancelRedirect,
		SuccessCallback:        request.SuccessCallback,
		PaymentMethodWhitelist: request.PaymentMethodWhitelist,
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Secret))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, err.Error())
		}
		errorMessage := string(body)
		return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, errorMessage)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response PurchaseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientConfig) RetrievePurchase(id string) (*PurchaseResponse, error) {
	url := BASE_URL + "/purchases/" + id + "/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Secret))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, err.Error())
		}
		errorMessage := string(body)
		return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, errorMessage)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response PurchaseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientConfig) RefundPurchase(id string) (*RefundPurchaseResponse, error) {
	url := BASE_URL + "/purchases/" + id + "/refund/"

	// payload, err := json.Marshal(map[string]interface{}{"amount": 120})
	// if err != nil {
	// 	return nil, err
	// }

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Secret))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, err.Error())
		}
		errorMessage := string(body)
		return nil, fmt.Errorf("status code: %d, error: %s", res.StatusCode, errorMessage)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response RefundPurchaseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
