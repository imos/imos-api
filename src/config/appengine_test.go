// +build !appengine

package config_test

import (
	"config"
	"github.com/imos/imosrpc"
	"reflect"
	"testing"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestAppEngine(t *testing.T) {
	response := config.AppEngine()
	expectedResponse := config.AppEngineResponse{
		IsAppEngine: false,
		IsProduction: false,
	}
	if !reflect.DeepEqual(expectedResponse, response) {
		t.Errorf("expected: %#v, actual: %#v.", expectedResponse, response)
	}
}
