package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//ログファイル作成
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//エラーハンドリング
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
