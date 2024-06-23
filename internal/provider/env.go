package provider

import (
	"fmt"
	"os"
	"strings"

	env "github.com/ProovGroup/lib-env"
)

var e env.Env

func init() {
	if os.Getenv("TEST") == "true" {
		return
	}
	bucketEnvName := "weproov-environment-" + os.Getenv("ENV") + "-files"
	files := []string{
		"s3://" + bucketEnvName + "/V1/env_core_main_database.json:eu-west-1",
		"s3://" + bucketEnvName + "/V1/env_claim_database.json:eu-west-1",
	}

	ev, err := env.NewEnvironment(strings.Join(files, ","))

	if err != nil {
		fmt.Println("[ERROR] NewEnvironment:", err)
		panic(err) // App cannot work without proper env
	}

	e = ev
}

func GetEnv() env.Env {
	return e
}
