package main

import (
	"github.com/universalmacro/auth/controllers"
	"github.com/universalmacro/common/config"
)

func main() {
	controllers.Init(":" + config.GetString("server.port"))
}
