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

	tickerChannel := make(chan bitflyer.Ticker)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
	for ticker := range tickerChannel {
		fmt.Println(ticker)
		fmt.Println(ticker.GetMidPrice())
		fmt.Println(ticker.DateTime())
		fmt.Println(ticker.TruncateDateTime(time.Second))
		fmt.Println(ticker.TruncateDateTime(time.Minute))
		fmt.Println(ticker.TruncateDateTime(time.Hour))
	}
}
