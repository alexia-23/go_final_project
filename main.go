package main

import (
	"log"

	"github.com/alexia-23/go_final_project/pkg/database"
	"github.com/alexia-23/go_final_project/pkg/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env:", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatal("Не удалось настроить базу:", err)
	}

	s := server.New(db)
	defer s.DB.Close()
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
	err = s.Run()
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
