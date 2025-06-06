package inbound

import "github.com/jeffotoni/quick"

type HTTPHandler interface {
	RegisterRoutes(router *quick.App)
}
