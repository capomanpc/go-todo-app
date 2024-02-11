package config

import (
	"log"

	"github.com/capomanpc/go-todo-app/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

// main関数より前に実行
func init() {
	LoadConfig()
	utils.LoggingSettings()
}

func LoadConfig() {
	cfg, err := ini.Load("coonfig.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("Port").MustString("8080"),
		SQLDriver: cfg.Section("web").Key("SQLDriver").String(),
		DBName:    cfg.Section("web").Key("DBName").String(),
		LogFile:   cfg.Section("web").Key("Logfile").String(),
	}
}
