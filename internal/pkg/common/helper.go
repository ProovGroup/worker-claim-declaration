package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ProovGroup/lib-claim-models/prequalif"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/common/model"
	"github.com/ProovGroup/worker-claim-declaration/internal/provider"
)

// Create a standard sinister struct containing informations about
// the sinister to be declared using the data available in the
// files and prequalif tables
func GetSinister(proovCode string) (*model.Sinister, error) {
	// Get prequalif from db
	pr, err := prequalif.GetPrequalifByProovCode(provider.GetClaimDB(), proovCode)
	if err != nil {
		fmt.Println("[ERROR] prequalif.GetPrequalifById:", err)
		return nil, err
	}

	// Get sinister type from JsonModel within the prequalif struct
	var jsonModel map[string]interface{}
	if err = json.Unmarshal([]byte(pr.JsonModel), &jsonModel); err != nil {
		fmt.Printf("[ERROR] json_model is not valid JSON (prequalif id: %d)\n", pr.ID)
		return nil, err
	}

	return &model.Sinister{
		Prequalif: &pr,
		ProovCode: proovCode,
		JsonModel: jsonModel,
		Register:  pr.Register,
		CreatedAt: *pr.CreatedAt,
		UpdatedAt: *pr.UpdatedAt,
	}, nil
}

// Set-up and send a POST request to a URL with a given body and access token
// and returns the response as a byte array
func PostSinister(url string, body []byte) ([]byte, error) {
	token, err := provider.GetToken()
	if err != nil {
		fmt.Println("[ERROR] provider.GetToken:", err)
		return nil, err
	}

	// Initialize request
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("[ERROR] http.NewRequest:", err)
		return nil, err
	}

	// Add token to request
	req.Header.Add("Authorization", token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR] client.Do:", err)
		return nil, err
	}

	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[ERROR] io.ReadAll:", err)
		return nil, err
	}

	if !strings.HasPrefix(resp.Status, "2") {
		fmt.Println("[ERROR] Status code != 2XX:", resp.Status, string(respBody))
		return nil, fmt.Errorf("%s: %s", resp.Status, string(respBody))
	}

	return respBody, nil
}
