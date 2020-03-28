package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/RainrainWu/probe/pkg/test_salad"
	"github.com/RainrainWu/probe/pkg/utils"
)

const (
	SERVICE_PORT	= "2023"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func testHandler(ctx *gin.Context) {

	var meta utils.Metadata
	err := ctx.BindJSON(&meta)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	salad_runner := utils.Runner{
		Series: 	test_salad.Cases,
		LogLevel:	1,
	}
	salad_runner.Init()
	salad_runner.Rep.SetMeta(meta)
	go salad_runner.Start()
	ctx.JSON(200, []string{"123", "321"})
}

func main() {

	server := gin.Default()
	server.POST("/test", testHandler)
	server.Run("localhost:" + SERVICE_PORT)
}
