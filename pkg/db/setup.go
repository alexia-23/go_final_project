package db

import (
	"database/sql"
	"io/ioutil"
	"os"

	"github.com/alexia-23/go_final_project/pkg/logger"
	_ "github.com/mattn/go-sqlite3"
)

func Setup() {
	log := logger.Get()
	dataSource := os.Getenv("DATA_SOURCE")
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	script, err := ioutil.ReadFile("sql/v1.sql")
	if err != nil {
		log.Fatal("не удалось прочитать файл миграции:", err)
	}

	if _, err := db.Exec(string(script)); err != nil {
		log.Fatal("ошибка при выполнении миграции:", err)
	}

	log.Println("Миграция успешно выполнена!")
}
