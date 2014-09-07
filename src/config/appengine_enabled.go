// +build appengine

package config

import (
	"appengine_internal"
)

func IsAppEngine() bool {
	return true
}

func IsProduction() bool {
	return !appengine_internal.IsDevAppServer()
}
