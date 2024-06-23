package pkg

import (
	"fmt"
	"strings"

	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/adhoc"
	"github.com/ProovGroup/worker-claim-declaration/internal/pkg/veos"
)

type API string

type DeclarationParams struct {
	Api string `json:"api"`
}

type Payload struct {
	ProovCode string            `json:"proov_code"`
	Params    DeclarationParams `json:"params"`
}

func Handler(payload Payload) error {
	if payload.ProovCode == "" {
		panic("Missing proov_code in payload")
	}

	var err error
	switch strings.ToLower(string(payload.Params.Api)) {
	case "veos":
		err = veos.Handle(payload.ProovCode)
	case "ad-hoc":
		fallthrough
	case "adhoc":
		err = adhoc.Handle(payload.ProovCode)
	default:
		fmt.Printf("'%s' is not a valid API name\n", string(payload.Params.Api))
	}

	if err != nil {
		fmt.Println("[ERROR] Handle error:", err)
		return err
	}

	return nil
}
