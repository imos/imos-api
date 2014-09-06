package init

import (
	"github.com/imos/imosrpc"
)

func init() {
  http.HandleFunc("/", imosrpc.DefaultHandler())
}
