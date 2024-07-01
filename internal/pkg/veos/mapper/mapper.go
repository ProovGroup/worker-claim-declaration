package mapper

import (
	"fmt"
	"time"

	commonModel "github.com/ProovGroup/worker-claim-declaration/internal/pkg/common/model"
	veosModel "github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos/model"
)

// Transforms a common Sinister struct into a veos Sinister struct
func MapSinisterToVEOS(sinister *commonModel.Sinister) *veosModel.Sinister {
	var ok bool

	// Used to get the time in France as it is not stored with timezone information
	loc, _ := time.LoadLocation("Europe/Paris")

	// Set the sinister type to either material (1) or corporal (2)
	var isCorporalDamage bool
	if isCorporalDamage, ok = sinister.JsonModel["corporalDamage"].(bool); !ok {
		fmt.Println("[ERROR] missing 'corporalDamage' field in sinister json_model")
		return nil
	}
	sinisterType := "1"
	if isCorporalDamage {
		sinisterType = "2"
	}

	// Get the id_pol_survenance from contractInfo within the json_model
	var contractInfo veosModel.ContractInfo
	if contractInfo, ok = sinister.JsonModel["contractInfo"].(veosModel.ContractInfo); !ok {
		fmt.Println("[ERROR] missing 'contractInfo' field in sinister json_model")
		return nil
	}

	return &veosModel.Sinister{
		IDPol:           contractInfo.IdPolSurvenance,
		NumCbt:          sinister.ProovCode,
		Type:            sinisterType,
		DateOuverture:   sinister.CreatedAt.Format("02/01/2006"),
		DateSurvenance:  sinister.CreatedAt.Format("02/01/2006"),
		HeureSurvenance: sinister.CreatedAt.In(loc).Format(time.TimeOnly),
		DateDeclaration: sinister.UpdatedAt.Format("02/01/2006"),
	}
}
