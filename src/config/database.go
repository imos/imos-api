package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/imos/imosql"
)

var connection *imosql.Connection
var targetName = "None"

func init() {
	// Connect to the SQL server first.
	GetDatabase()
}

func GetDatabase() *imosql.Connection {
	if connection != nil {
		return connection
	}
	var err error
	if IsAppEngine() && IsProduction() {
		connection, err = imosql.Open(
			"mysql", "api@cloudsql(imos-api:sql)/api?timeout=5s")
		if err == nil {
			targetName = "Cloud SQL"
			return connection
		}
	}
	if IsAppEngine() && !IsProduction() {
		connection, err = imosql.Open(
			"mysql", "api@tcp(172.17.42.1:3306)/api?timeout=5s")
		if err == nil {
			targetName = "MySQL"
			return connection
		}
	}
	return nil
}

type DatabaseInfo struct {
	TargetName string
}

func GetDatabaseInfo() DatabaseInfo {
	return DatabaseInfo{
		TargetName: targetName,
	}
}
