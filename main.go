package main

import (
	"log"

	"github.com/junichiseki0831/gotrading/config"
	"github.com/junichiseki0831/gotrading/utils"
)

func main() {
	// configから読み取った内容を表示
	//fmt.Println(config.Config.ApiKey)
	//fmt.Println(config.Config.ApiSecret)
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
