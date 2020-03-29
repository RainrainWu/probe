package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RainrainWu/probe/pkg/test_salad"
	"github.com/RainrainWu/probe/pkg/utils"
)

const (
	SERVICE_PORT	= "2023"
)

func testHandler(ctx *gin.Context) {

	var meta utils.Metadata
	err := ctx.BindJSON(&meta)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	salad_runner := utils.Runner{
		Series: 	test_salad.Cases,
	}
	salad_runner.Init()
	salad_runner.Rep.SetMeta(meta)
	go salad_runner.Run()
	ctx.Data(200, "plain/text", []byte(<- salad_runner.Result))
}

func main() {

	server := gin.Default()
	server.POST("/test", testHandler)
	server.Run("localhost:" + SERVICE_PORT)
}
