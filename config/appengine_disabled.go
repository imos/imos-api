// +build !appengine

package config

func IsAppEngine() bool {
	return false
}

func IsProduction() bool {
	return false
}
