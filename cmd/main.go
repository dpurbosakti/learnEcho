package main

import (
	"learn-echo/config"
	"learn-echo/internal/handlers"
)

func main() {
	config.Run(handlers.SetRoutes(), "test")
}
