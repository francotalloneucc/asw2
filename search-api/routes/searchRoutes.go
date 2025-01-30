package routes

import (
	"search-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/hotels/index", controllers.IndexHotel)
	return r
}
