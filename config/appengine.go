package config

import (
	"github.com/imos/imosrpc"
)

func init() {
	imosrpc.RegisterHandler("config/appengine", DatabaseHandler)
}

type AppEngineRequest struct {}
type AppEngineResponse struct {
	IsAppEngine bool `json:"is_app_engine"`
	IsProduction bool `json:"is_production"`
}

func AppEngineHandler(request AppEngineRequest) AppEngineResponse {
	return AppEngineResponse{
		IsAppEngine: IsAppEngine(),
		IsProduction: IsProduction(),
	}
}

func AppEngine() AppEngineResponse {
	request := AppEngineRequest{}
	response := AppEngineResponse{}
	imosrpc.Call("config/appengine", request, &response)
	return response
}
