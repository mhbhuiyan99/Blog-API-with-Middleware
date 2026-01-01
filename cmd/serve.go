package cmd

import (
	"blogAPI/config"
	"blogAPI/rest"
)

func Serve() {

	cnf :=config.GetConfig() // Load configuration

	rest.Start(cnf) // Start the REST server
	
}