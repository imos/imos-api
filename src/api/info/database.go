package info

import (
	"config"
	"github.com/imos/imosrpc"
)

func init() {
	imosrpc.RegisterHandler("info/database", DatabaseHandler)
}

type DatabaseRequest struct{}
type DatabaseResponse struct {
	config.DatabaseInfo
}

func DatabaseHandler(request DatabaseRequest) DatabaseResponse {
	return DatabaseResponse{DatabaseInfo: config.GetDatabaseInfo()}
}

func Database() DatabaseResponse {
	request := DatabaseRequest{}
	response := DatabaseResponse{}
	imosrpc.Call("info/database", request, &response)
	return response
}
