package probe

import (
	"fmt"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RainrainWu/probe/jobs"
	"github.com/RainrainWu/probe/utils"
	"github.com/RainrainWu/probe/config"
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

	var filter utils.Filter
	err := ctx.BindJSON(&filter)
	if err != nil {
		utils.Logger.Warn("Failed to bind json data")
		ctx.Data(400, "plain/text", []byte("Unacceptable json structure"))
	}

	result := utils.ReadReportRaw(filter.Index)
	text, _ := utils.Render(result)
	ctx.Data(200, "plain/text", []byte(text))
}

// handle metrix response
func metrixHandler(ctx *gin.Context) {

	role, _ := ctx.Get("role")
	if role != "UDC Tester" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var filter utils.Filter
	err := ctx.BindJSON(&filter)
	if err != nil {
		utils.Logger.Warn("Failed to bind json data")
		ctx.Data(400, "plain/text", []byte("Unacceptable json structure"))
	}

	result := utils.ReadReport(filter.Index)
	ctx.Data(200, "plain/text", []byte(result))
}

func Start() {

	server := gin.Default()
	server.POST("/login", loginHandler)
	server.GET("/report", reportHandler)
	server.GET("/report/metrix", metrixHandler)

	// Auth required endpoint
	authorized := server.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.GET("/foo", fooHandler)
		authorized.POST("/test", testHandler)
	}

	server.Run("localhost:" + config.SERVICE_PORT)
}
