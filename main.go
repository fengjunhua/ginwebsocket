package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsupgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter,r *http.Request)  {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed to webSocket upGrader")
		return
	}
	for{
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(msg)
		conn.WriteMessage(t,[]byte("pong"))
	}

}

func main() {
	router := gin.Default()
    router.LoadHTMLFiles("views/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)

	})
	//webSocket 请求使用 wsHandler函数处理
    router.GET("/ws", func(c *gin.Context) {
    	wsHandler(c.Writer,c.Request)
	})
	

	router.Run(":12312")
}
