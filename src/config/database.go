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
	user := "api"
	if !ShouldUseProductionDatabase() {
		user = "api-experimental"
	}
	if IsAppEngine() && IsProduction() {
		connection, err = imosql.Open(
			"mysql", user+"@cloudsql(imos-api:sql)/"+user+"?timeout=5s")
		if err == nil {
			targetName = "Cloud SQL"
			return connection
		}
	}
	if IsAppEngine() && !IsProduction() {
		connection, err = imosql.Open(
			"mysql", user+"@tcp(173.194.111.104:3306)/"+user+"?timeout=5s")
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
