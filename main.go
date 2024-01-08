package main

import (
	"login-ports/config"
	"login-ports/lib/env"
)

func main() {
	cfg := config.NewConfig()

	cfg.Router.Run(env.String("MainSetup.ServerHost", "3000"))
}
