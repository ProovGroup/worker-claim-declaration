package veos

import (
	"encoding/json"
	"fmt"

	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/common"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos/mapper"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos/model"
	"github.com/ProovGroup/worker-claim-declaration/internal/provider"
)

var (
	BASE_URL          = provider.Getenv(provider.API_VEOS_URL)
	ROUTE_DECLARATION = BASE_URL + "/sinistre"
)

func Handle(proovCode string) error {
	// Get common sinister from prequalif table infos
	sinister, err := common.GetSinister(proovCode)
	if err != nil {
		fmt.Println("[ERROR] common.GetSinister:", err)
		return err
	}

	// Create request body
	requestJSON, err := json.Marshal(mapper.MapSinisterToVEOS(sinister))
	if err != nil {
		fmt.Println("[ERROR] json.Marshal:", err)
		return err
	}

	// Send sinister to VEOS
	respBody, err := common.PostSinister(ROUTE_DECLARATION, requestJSON)
	if err != nil {
		fmt.Println("[ERROR] common.PostSinister:", err)
		return err
	}

	// Parse response
	var respSinister model.ResponseSinister
	if err := json.Unmarshal(respBody, &respSinister); err != nil {
		fmt.Println("[ERROR] json.Unmarshal:", err)
		return err
	}

	// Check response is ok
	if respSinister.Id == -1 {
		fmt.Println("[ERROR] Sinister not created:", respSinister.Message)
		return fmt.Errorf(respSinister.Message)
	}

	// Save response body in response_model field
	sinister.Prequalif.ResponseModel = string(respBody)
	err = sinister.Prequalif.Save(provider.GetClaimDB())
	if err != nil {
		fmt.Println("[ERROR] sinister.Prequalif.Save:", err)
		return err
	}

	fmt.Printf("[INFO][VEOS] Declaration done (proov_code: %s, sinister_id: %d)\n", proovCode, respSinister.Id)

	return nil
}
