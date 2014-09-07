package info

import (
	"config"
	"github.com/imos/imosrpc"
)

func init() {
	imosrpc.RegisterHandler("info/appengine", AppEngineHandler)
}

type AppEngineRequest struct{}
type AppEngineResponse struct {
	IsAppEngine  bool `json:"is_app_engine"`
	IsProduction bool `json:"is_production"`
}

func AppEngineHandler(request AppEngineRequest) AppEngineResponse {
	return AppEngineResponse{
		IsAppEngine:  config.IsAppEngine(),
		IsProduction: config.IsProduction(),
	}
}

func AppEngine() AppEngineResponse {
	request := AppEngineRequest{}
	response := AppEngineResponse{}
	imosrpc.Call("info/appengine", request, &response)
	return response
}
