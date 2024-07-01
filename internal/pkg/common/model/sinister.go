package model

import (
	"time"

	"github.com/ProovGroup/lib-claim-models/prequalif"
)

type Sinister struct {
	Prequalif *prequalif.Prequalif

	ProovCode  string
	JsonModel  map[string]interface{}
	Register   string
	IsCorporal bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
