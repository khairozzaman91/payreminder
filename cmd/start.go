package cmd

import (
	"payreminder/config"
	"payreminder/rest"
)

func Start() {
	cnf := config.GetConfig()
	rest.Server(&cnf)
}
