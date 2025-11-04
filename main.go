package main

import (
	"github.com/alexia-23/go_final_project/pkg/db"
	"github.com/alexia-23/go_final_project/pkg/logger"
	"github.com/alexia-23/go_final_project/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	logger.Init()
	log := logger.Get()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env:", err)
	}

	db.Setup()

	err = server.Init()
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}

}
