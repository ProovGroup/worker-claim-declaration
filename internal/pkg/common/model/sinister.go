package model

import "time"

type Sinister struct {
	ProovCode  string
	IDPol      string
	Register   string
	IsCorporal bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
