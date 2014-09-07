package infoapi

import (
	"github.com/imos/imosrpc"
	"time"
)

func init() {
	imosrpc.RegisterHandler("info/time", TimeHandler)
}

type TimeRequest struct{}

type TimeResponse struct {
	CurrentTime time.Time `json:"current_time"`
}

func TimeHandler(request TimeRequest) TimeResponse {
	return TimeResponse{CurrentTime: time.Now()}
}

func Time() time.Time {
	request := TimeRequest{}
	response := TimeResponse{}
	imosrpc.Call("info/time", request, &response)
	return response.CurrentTime
}
