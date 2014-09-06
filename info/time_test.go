package infoapi_test

import (
	"github.com/imos/imosrpc"
	"github.com/imos/infoapi"
	"testing"
	"time"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestTime(t *testing.T) {
	currentTime := time.Now()
	apiTime := infoapi.Time()
	if apiTime.Before(currentTime) || apiTime.After(time.Now()) {
		t.Errorf("Time returns a wrong time: %s.", apiTime)
	}
}
