package probe

import (
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RainrainWu/probe/utils"
	"github.com/RainrainWu/probe/config"
	"github.com/RainrainWu/probe/handlers"
)

// Authentic required middleware
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

func Start() {

	server := gin.Default()
	server.POST("/login", handlers.LoginHandler)
	server.GET("/report", handlers.ReportHandler)
	server.GET("/report/metrix", handlers.MetrixHandler)

	// Auth required endpoint
	authorized := server.Group("/")
	authorized.Use(AuthRequired)
	{
		authorized.POST("/test", handlers.ExecHandler)
	}

	server.Run("localhost:" + config.SERVICE_PORT)
}
