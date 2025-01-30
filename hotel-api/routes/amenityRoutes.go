package routes

import (
	"hotel-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAmenityRoutes(r *gin.Engine) {
	amenityController := controllers.AmenityController{}

	// Ruta para crear un nuevo amenity
	r.POST("/createAmenity", amenityController.CreateAmenity)

	// Ruta para obtener un amenity por ID
	r.GET("/getAmenityByID/:id", amenityController.GetAmenity)

	// Ruta para obtener todos los amenities
	r.GET("/getAllAmenities", amenityController.GetAmenities)

	// Ruta para actualizar un amenity
	r.PUT("/updateAmenity/:id", amenityController.UpdateAmenity)

	// Ruta para eliminar un amenity
	r.DELETE("/deleteAmenity/:id", amenityController.DeleteAmenity)
}
