package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web")) // раздаём статические файлы из ./web
	http.Handle("/", fs)

	port := ":7540"
	fmt.Println("🚀 Сервер запущен на http://localhost" + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
