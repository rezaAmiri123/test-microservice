package agent

import (
	"net/http"
)

func (a *Agent) setupKeepAlive() error {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	go http.ListenAndServe(a.HttpKeepAliveServerHostPort, nil)
	return nil
}
