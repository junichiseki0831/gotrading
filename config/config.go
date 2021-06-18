package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini") //config.iniファイルの読み込み

	// configファイルを読み込めなければログを出して抜ける
	if err != nil {
		log.Printf("Faild to read file: %v", err)
		os.Exit(1)
	}

	// ConfigListにconfig.iniの内容をstringで格納する
	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:   cfg.Section("gotrading").Key("log_file").String(),
	}
}
