package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/RainrainWu/probe/jobs"
	"github.com/RainrainWu/probe/utils"
)

// handle login require
func LoginHandler(ctx *gin.Context) {

	user := utils.NewUser()
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if (user.CheckUser()) {
		utils.Logger.Info("User Login",
			zap.String("User", user.Username),
		)
		ctx.JSON(http.StatusOK, gin.H{
			"token": user.GenToken(),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect username or password",
		})
	}
}


// handle test executing
func ExecHandler(ctx *gin.Context) {

	// Check authentication
	role, ok := ctx.Get("role")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get user role from request",
		})
	} else if role != "UDC Tester" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Check metadata description
	var meta utils.Metadata
	err := ctx.BindJSON(&meta)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Run jobs
	utils.Logger.Info("Test Execute",
		zap.String("Index", meta.Index),
		zap.String("Tester", meta.Tester),
	)
	result := jobs.RunJob(meta)
	ctx.Data(http.StatusOK, "plain/text", []byte(result))
}

// handle report viewing
func ReportHandler(ctx *gin.Context) {

	var filter utils.Filter
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		utils.Logger.Warn("Failed to bind json data")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unknown json structure",
		})
		return
	}

	utils.Logger.Info("Read Report",
		zap.String("Index", filter.Index),
	)
	result := utils.ReadReportRaw(filter.Index)
	text, _ := utils.Render(result)
	ctx.Data(200, "plain/text", []byte(text))
}


// handle metrix response
func MetrixHandler(ctx *gin.Context) {

	var filter utils.Filter
	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		utils.Logger.Warn("Failed to bind json data")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unknown json structure",
		})
		return
	}

	utils.Logger.Info("Read Report",
		zap.String("Index", filter.Index),
	)
	result := utils.ReadReport(filter.Index)
	ctx.Data(http.StatusOK, "plain/text", []byte(result))
}