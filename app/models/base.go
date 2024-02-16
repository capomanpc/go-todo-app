package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/capomanpc/go-todo-app/config"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// ユーザーモデルの作成

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	// 自動増分、uuidは一意になるように設定
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name string,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		content text,
		user_id integer,
		created_at datetime)`, tableNameTodo)

	Db.Exec(cmdT)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
