package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ProovGroup/lib-claim-models/files"
	"github.com/ProovGroup/lib-claim-models/prequalif"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/common/model"
	"github.com/ProovGroup/worker-claim-declaration/internal/provider"
)

// Create a standard sinister struct containing informations about
// the sinister to be declared using the data available in the
// files and prequalif tables
func GetSinister(proovCode string) (*model.Sinister, error) {
	// Get files from db
	f, err := files.GetFilesByProovCode(provider.GetClaimDB(), proovCode)
	if err != nil {
		fmt.Println("[ERROR] files.GetFilesByProovCode:", err)
		return nil, err
	}

	// Get prequalif from db
	pr, err := prequalif.GetPrequalifById(provider.GetClaimDB(), f.PrequalifId)
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

	var (
		isCorporalDamage bool
		ok               bool
	)
	if isCorporalDamage, ok = jsonModel["corporalDamage"].(bool); !ok {
		msg := fmt.Sprintf("corporalDamage field is missing in the prequalif's JsonModel or it is not a valid boolean value (prequalif id: %d)", pr.ID)
		fmt.Println("[ERROR] " + msg)
		return nil, fmt.Errorf(msg)
	}

	return &model.Sinister{
		ProovCode:  proovCode,
		IDPol:      f.IdPol,
		Register:   pr.Register,
		IsCorporal: isCorporalDamage,
		CreatedAt:  *f.CreatedAt,
		UpdatedAt:  *f.UpdatedAt,
	}, nil
}

// Used once the target API has responded with a sinister_id that we should store in the files
func SaveSinisterId(proovCode string, sinisterId int) error {
	f, err := files.GetFilesByProovCode(provider.GetClaimDB(), proovCode)
	if err != nil {
		fmt.Println("[ERROR] files.GetFilesByProovCode:", err)
		return err
	}

	f.SinistreId = sinisterId
	return f.Save(provider.GetClaimDB())
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
