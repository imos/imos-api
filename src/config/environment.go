package config

import "environment"

func IsMaster() bool {
	return IsProduction() && environment.Branch == "master"
}
