package mapper

import (
	"time"

	commonModel "github.com/ProovGroup/worker-claim-declaration/internal/pkg/common/model"
	veosModel "github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos/model"
)

// Transforms a common Sinister struct into a veos Sinister struct
func MapSinisterToVEOS(sinister *commonModel.Sinister) *veosModel.Sinister {
	// Used to get the time in France as it is not stored with timezone information
	loc, _ := time.LoadLocation("Europe/Paris")

	// Set the sinister type to either material (1) or corporal (2)
	sinisterType := "1"
	if sinister.IsCorporal {
		sinisterType = "2"
	}

	return &veosModel.Sinister{
		IDPol:           sinister.IDPol,
		NumCbt:          sinister.ProovCode,
		Type:            sinisterType,
		DateOuverture:   sinister.CreatedAt.Format("02/01/2006"),
		DateSurvenance:  sinister.CreatedAt.Format("02/01/2006"),
		HeureSurvenance: sinister.CreatedAt.In(loc).Format(time.TimeOnly),
		DateDeclaration: sinister.UpdatedAt.Format("02/01/2006"),
	}
}
