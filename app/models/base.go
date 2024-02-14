package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/capomanpc/go-todo-app/config"
	_ "github.com/mattn/go-sqlite3"
)

// ユーザーモデルの作成

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	// 自動増分、uuidは一意になるように設定
	cmdU := fmt.Sprintf(`CERATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMAARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}
