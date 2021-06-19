package main

import (
	"fmt"

	"github.com/junichiseki0831/gotrading/bitflyer"
	"github.com/junichiseki0831/gotrading/config"
	"github.com/junichiseki0831/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile) //ログ作成
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
