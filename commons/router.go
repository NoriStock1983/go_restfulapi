package commons

import (
	"go_restfulapi/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	mastermaintenancerouter := router.Group("/master")
	{
		usermst := mastermaintenancerouter.Group("/usermst")
		{
			usermst.GET("/allusers", controllers.AllUsers)
		}
	}
	return router
}
