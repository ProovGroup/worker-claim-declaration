package provider

import (
	"fmt"
	"os"
)

type ENV_VAR string

var (
	API_VEOS_URL  ENV_VAR = "API_VEOS_URL"
	API_ADHOC_URL ENV_VAR = "API_ADHOC_URL"
)

var requiredVariables = []ENV_VAR{
	API_VEOS_URL,
	API_ADHOC_URL,
}

func init() {
	// Make sure all the required environment variables are set
	for _, envVar := range requiredVariables {
		if os.Getenv(string(envVar)) == "" {
			panic(fmt.Sprintf("Missing '%s' environment variable!", envVar))
		}
	}
}

func Getenv(v ENV_VAR) string {
	return os.Getenv(string(v))
}
