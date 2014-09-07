// +build appengine

package config

import (
	"appengine_internal"
	_ "github.com/go-sql-driver/mysql"
	"github.com/imos/imosql"
	"github.com/imos/imosrpc"
)

func IsAppEngine() bool {
	return true
}

func IsProduction() bool {
	return !appengine_internal.IsDevAppServer()
}
