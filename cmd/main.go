package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/RainrainWu/probe/pkg/jobs"
	"github.com/RainrainWu/probe/pkg/utils"
)

const (
	SERVICE_PORT string	= "2023"
)

func testHandler(ctx *gin.Context) {

	var meta utils.Metadata
	err := ctx.BindJSON(&meta)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	result := jobs.RunJob(meta)
	ctx.Data(200, "plain/text", []byte(result))
}

func reportHandler(ctx *gin.Context) {

	var filter utils.Filter
	err := ctx.BindJSON(&filter)
	utils.HandleErr(err, "Failed to bind json data")

	result := utils.ReadReport(filter.Index)
	ctx.Data(200, "plain/text", []byte(result))
}

func main() {

	server := gin.Default()
	server.POST("/test", testHandler)
	server.GET("/report", reportHandler)
	server.Run("localhost:" + SERVICE_PORT)
}
