// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/eeianux/lmq/middleware"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.Use(middleware.Logger())

	// your code ...
}