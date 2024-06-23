package model

type Sinister struct {
	IDPol           string `json:"idPol"`
	NumCbt          string `json:"numCbt"`
	Type            string `json:"type"`
	DateOuverture   string `json:"dateOuverture"`
	DateSurvenance  string `json:"dateSurvenance"`
	HeureSurvenance string `json:"heureSurvenance,omitempty"`
	DateDeclaration string `json:"dateDeclaration"`
}
