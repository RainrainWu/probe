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

func main() {

	server := gin.Default()
	server.POST("/test", func(c *gin.Context) {
		salad_runner := utils.Runner{
			Series: 	test_salad.Cases,
			LogLevel:	1,
		}
		salad_runner.Init("test_001", "dev", "rain", "first test")
		go salad_runner.Start()
		show(salad_runner)
		c.JSON(200, []string{"123", "321"})
	})
	server.Run("localhost:" + SERVICE_PORT)
}

func show(runner utils.Runner) {

	for {
		select {
		case msg := <- runner.Logger:
			fmt.Println(msg)
		default:
		}
	}
}