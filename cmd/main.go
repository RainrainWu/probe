package main

import (
	"fmt"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RainrainWu/probe/pkg/jobs"
	"github.com/RainrainWu/probe/pkg/utils"
	"github.com/RainrainWu/probe/pkg/config"
)

// handle login require
func loginHandler(ctx *gin.Context) {

	user := utils.NewUser()
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if (user.CheckUser()) {
		ctx.JSON(http.StatusOK, gin.H{
			"token": user.GenToken(),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect username or password",
		})
	}
}

func AuthRequired(ctx *gin.Context) {

	auth := ctx.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]
	message, tokenClaims := utils.ValidateToken(token)
	if (message != "") {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		ctx.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*utils.Claims); ok && tokenClaims.Valid {
		ctx.Set("user", claims.User)
		ctx.Set("role", claims.Role)
		ctx.Next()
	} else {
		ctx.Abort()
		return
	}
}

func fooHandler(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"token": ctx.GetHeader("Authorization"),
	})
}

// handle test executing
func testHandler(ctx *gin.Context) {

	role, _ := ctx.Get("role")
	if role != "UDC Tester" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var meta utils.Metadata
	err := ctx.BindJSON(&meta)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	result := jobs.RunJob(meta)
	ctx.Data(200, "plain/text", []byte(result))
}

// handle report viewing
func reportHandler(ctx *gin.Context) {

	role, _ := ctx.Get("role")
	if role != "UDC Tester" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var filter utils.Filter
	err := ctx.BindJSON(&filter)
	utils.HandleErr(err, "Failed to bind json data")

	result := utils.ReadReport(filter.Index)
	ctx.Data(200, "plain/text", []byte(result))
}

func main() {

	server := gin.Default()
	server.POST("/login", loginHandler)

	// Auth required endpoint
	authorized := server.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.GET("/foo", fooHandler)
		authorized.POST("/test", testHandler)
		authorized.GET("/report", reportHandler)
	}

	server.Run("localhost:" + config.SERVICE_PORT)
}
