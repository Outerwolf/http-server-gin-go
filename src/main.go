package main

import (
	"github.com/outerwolf/http-server-gin-go/src/boot"
	"go.uber.org/fx"
)

func main() {
	println("Hello, world!")
	newFxApp().Run()
}

func newFxApp() *fx.App {
	return fx.New(
		fx.Provide(
			boot.CreateHTTPGinServer,
		),
		fx.Invoke(
			boot.RegisterHttpRoutesGin,
			boot.StartHttpGinServer,
		),
	)
}
