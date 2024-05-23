package main

import "github.com/TiagoDiass/fullcycle-golang-rest-api/configs"

func main() {
	cfg, _ := configs.LoadConfig(".")

	println(cfg.DBName)
}
