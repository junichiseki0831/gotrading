package main

import (
	"fmt"
	"time"

	"github.com/junichiseki0831/gotrading/bitflyer"
	"github.com/junichiseki0831/gotrading/config"
	"github.com/junichiseki0831/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile) //ログ作成
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.DateTime())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}
