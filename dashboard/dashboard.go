package dashboard

import (
	"github.com/blockpilabs/near-drpc/network/http"
)

func ListenAndServ(router *Router) error {
	server := &http.HttpServer{
		Endpoint: ":8181",
		Handler:  router.router.Handler,

	}
	return server.ListenAndServe()
}
