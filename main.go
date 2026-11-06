package main

import (
	"payreminder/cmd"
)

// @title       PayReminder API
// @version     1.0
// @description PayReminder Backend API Documentation
// @host        localhost:8080
// @BasePath    /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	cmd.Start()
}
