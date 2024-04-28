package main

import (
	"tasktracker/internal/app"
)

const configsDir = "configs/"

//const (
//	envLocal = "local"
//	envDev   = "dev"
//)

func main() {
	app.Run(configsDir)
}
