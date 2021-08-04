// Package livez implements livez handlers.
package livez

import (
	"fmt"
	"net/http"

	"sigs.k8s.io/aws-encryption-provider/pkg/plugin"
)

// NewHandler returns a new livez handler.
func NewHandler(p *plugin.Plugin) http.Handler {
	return &handler{p: p}
}

type handler struct {
	p *plugin.Plugin
}

func (hd *handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := hd.p.Check()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, http.StatusText(http.StatusOK))
}
