package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"telegraph-search/src/search"

	"github.com/gin-gonic/gin"
)

func recv(channel chan search.SearchResult) {
	recv_count := 0

	for recv_count != 12*31*10 {
		result := <-channel

		log.Println(result)

		recv_count++
	}
}

func downloadFile(url string, filepath string) error {
	// Создаем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем файл для сохранения
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Копируем данные из HTTP-ответа в файл
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	r := gin.Default()

	r.Static("/static", "./web/static")

	r.GET("/", indexHandler)
	r.GET("/ws", wsHandler)

	r.Run(":8080")

}
