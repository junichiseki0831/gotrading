package main

import (
	"fmt"

	"github.com/junichiseki0831/gotrading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}