package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {

	// ファイルに書き込むときは必ず末尾に書き込むことを指定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// SetOutputでlogの書き込み先を二つに指定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
