package main

import (
	"github.com/sushi86/WAConfigSetter/config"
	"github.com/sushi86/WAConfigSetter/server"
)

func main() {

	if !config.IsValidConfigFile() {
		// config file is not valid, start config create server
		server.Start()
	}

}
