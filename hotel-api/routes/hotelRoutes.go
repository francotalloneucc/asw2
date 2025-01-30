package routes

import (
	"hotel-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupHotelRoutes(r *gin.Engine) {
	hotelController := &controllers.HotelController{}

	// Rutas para los hoteles con palabras clave
	r.POST("/hotels/createHotel", hotelController.CreateHotel)       // /createHotel: Crear un nuevo hotel
	r.GET("/hotels/getHotels", hotelController.GetHotels)            // /getHotels: Obtener todos los hoteles
	r.GET("/hotels/getHotel/:id", hotelController.GetHotel)          // /getHotel/:id: Obtener un hotel por ID
	r.PUT("/hotels/updateHotel/:id", hotelController.UpdateHotel)    // /updateHotel/:id: Actualizar un hotel
	r.DELETE("/hotels/deleteHotel/:id", hotelController.DeleteHotel) // /deleteHotel/:id: Eliminar un hotel
	r.GET("/hotels/check-existence/:hotelID", controllers.CheckHotelExistence)
}
