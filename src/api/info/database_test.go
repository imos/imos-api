package info_test

import (
	"config"
	"api/info"
	"github.com/imos/imosrpc"
	"reflect"
	"testing"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestDatabase(t *testing.T) {
	response := info.Database()
	expectedResponse := info.DatabaseResponse{
		DatabaseInfo: config.DatabaseInfo {
			TargetName: "None",
		},
	}
	if !reflect.DeepEqual(expectedResponse, response) {
		t.Errorf("expected: %#v, actual: %#v.", expectedResponse, response)
	}
}
