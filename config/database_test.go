package config_test

import (
	"github.com/imos/imos-api/config"
	"github.com/imos/imosrpc"
	"testing"
	"reflect"
)

func init() {
	imosrpc.SetHostname(imosrpc.InternalHostname)
}

func TestDatabase(t *testing.T) {
	response := config.Database()
	expectedResponse := config.DatabaseResponse{
		TargetName: "None",
	}
	if !reflect.DeepEqual(expectedResponse, response) {
		t.Errorf("expected: %#v, actual: %#v.", expectedResponse, response)
	}
}
