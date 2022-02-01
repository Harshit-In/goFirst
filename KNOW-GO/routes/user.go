package routes

import (
	"fastearn/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/user_info", controllers.GetAllBuyOrder())

}
