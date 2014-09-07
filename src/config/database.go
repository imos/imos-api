package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/imos/imosql"
	"github.com/imos/imosrpc"
)

var connection *imosql.Connection
var targetName = "None"

func init() {
	// Connect to the SQL server first.
	GetSql()
	imosrpc.RegisterHandler("config/database", DatabaseHandler)
}

func GetSql() *imosql.Connection {
	if connection != nil {
		return connection
	}
	var err error
	connection, err = imosql.Open(
		"mysql", "api@cloudsql(imos-api:sql)/api?timeout=5s")
	if err == nil {
		targetName = "Cloud SQL"
		return connection
	}
	connection, err := imosql.Open(
		"mysql", "api@tcp(172.17.42.1:3306)/api?timeout=5s")
	if err == nil {
		targetName = "MySQL"
		return connection
	}
	return nil
}

type DatabaseRequest struct{}
type DatabaseResponse struct {
	TargetName string `json:"target_name"`
}

func DatabaseHandler(request DatabaseRequest) DatabaseResponse {
	return DatabaseResponse{TargetName: targetName}
}

func Database() DatabaseResponse {
	request := DatabaseRequest{}
	response := DatabaseResponse{}
	imosrpc.Call("config/database", request, &response)
	return response
}
