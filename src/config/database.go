package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/imos/imosql"
)

var connection *imosql.Connection
var targetName = "None"

func init() {
	// Connect to the SQL server first.
	connection = connectDatabase()
}

func connectDatabase() *imosql.Connection {
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

func GetDatabase() *imosql.Connection {
	return connection
}

type DatabaseInfo struct {
	TargetName string `json:"target_name" sql:"target_name"`
	ConnectionId *int `json:"connection_id" sql:"connection_id"`
	User *string `json:"user" sql:"user"`
	Database *string `json:"database" sql:"database_name"`
	Version *string `json:"version" sql:"version"`
}

func GetDatabaseInfo() DatabaseInfo {
	if connection == nil {
		return DatabaseInfo{TargetName: "None"}
	}
	row := DatabaseInfo{}
	if !connection.RowOrDie(&row, `
		SELECT
			"MySQL" AS target_name,
			CONNECTION_ID() AS connection_id,
			USER() AS user,
			DATABASE() AS database_name,
			VERSION() AS version
	`) {
		return DatabaseInfo{}
	}
	return row
}
