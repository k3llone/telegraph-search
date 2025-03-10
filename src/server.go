package main

import (
	"encoding/json"
	"log"
	"net/http"
	"telegraph-search/src/search"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Разрешаем все соединения (для тестов)
	},
}

func indexHandler(ctx *gin.Context) {
	ctx.File("./web/index.html")
}

func wsHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to upgrade WebScoket"})
	}

	defer conn.Close()
	search_count := 0

	messageType, message, err := conn.ReadMessage()

	if err != nil {
		log.Println(err)
	}

	search_channel := make(chan search.SearchResult, 100)

	go search.RunSearch(string(message), search_channel)

	for {
		result := <-search_channel
		data, err := json.Marshal(result)

		if err != nil {
			log.Println(err)
		}

		if err = conn.WriteMessage(messageType, data); err != nil {
			log.Println(err)
			break
		}

		search_count += 1
	}

}
