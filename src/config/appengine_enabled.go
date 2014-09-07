// +build appengine

package config

import (
	"appengine"
)

func IsAppEngine() bool {
	return true
}

func IsProduction() bool {
	return !appengine.IsDevAppServer()
}
