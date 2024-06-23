package provider

import (
	env "github.com/ProovGroup/lib-env"
	"github.com/ProovGroup/lib-env/database"
)

var ClaimDB database.Database

func GetClaimDB() database.Database {
	if ClaimDB == nil {
		var success bool
		ClaimDB, success = e.GetDB(env.NewDBSelector("claim-db"))
		if !success {
			panic("Could not register the database claim-db")
		}
	}
	return ClaimDB
}
