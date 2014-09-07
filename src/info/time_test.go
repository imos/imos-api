package info_test

import (
	"github.com/imos/imosrpc"
	"info"
	"testing"
	"time"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestTime(t *testing.T) {
	currentTime := time.Now()
	apiTime := info.Time()
	if apiTime.Before(currentTime) || apiTime.After(time.Now()) {
		t.Errorf("Time returns a wrong time: %s.", apiTime)
	}
}
