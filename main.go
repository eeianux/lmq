// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/eeianux/lmq/middleware"
)

func main() {
	middleware.InitLog()
	h := server.Default()
	h.Use(middleware.Logger())
	register(h)
	h.Spin()
}
