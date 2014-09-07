package experimental

import (
	"appengine"
	"appengine/xmpp"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/experimental/send", sendHandler)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	m := &xmpp.Message{
		To:   []string{"i@imoz.jp"},
		Body: "This is a test.",
	}
	err := m.Send(c)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("failed to send: %s", err)))
		return
	}
	w.Write([]byte("a message is successfully sent."))
}
