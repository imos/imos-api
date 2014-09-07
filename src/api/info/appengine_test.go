// +build !appengine

package info_test

import (
	"api/info"
	"github.com/imos/imosrpc"
	"reflect"
	"testing"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestAppEngine(t *testing.T) {
	response := info.AppEngine()
	expectedResponse := info.AppEngineResponse{
		IsAppEngine: false,
		IsProduction: false,
	}
	if !reflect.DeepEqual(expectedResponse, response) {
		t.Errorf("expected: %#v, actual: %#v.", expectedResponse, response)
	}
}
