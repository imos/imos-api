package config

import (
	"environment"
	"fmt"
)

func init() {
	// Check if the branch name does not cause panic.
	ShouldUseProductionDatabase()
}

func IsMaster() bool {
	return IsProduction() && environment.Branch == "master"
}

func ShouldUseProductionDatabase() bool {
	switch environment.Branch {
	case "master":
		return true
	case "staging":
		return true
	case "devel":
		return false
	case "experimental":
		return false
	default:
		panic(fmt.Errorf("unknown branch: %s.", environment.Branch))
	}
}
