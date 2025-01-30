package controllers

import (
	"hotel-api/models"
	"hotel-api/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelController struct{}

// Crear un hotel
func (ctrl *HotelController) CreateHotel(c *gin.Context) {
	var hotelDto models.Hotel
	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := services.CreateHotel(hotelDto)
	if err != nil {
		// Si el error es de conflicto (hotel ya existe)
		if strings.Contains(err.Error(), "hotel with the same Name, Address, City, and Country already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ID": hotel.ID.Hex()}) // Usamos .Hex() para convertir ObjectID a string
}

// Obtener un hotel por ID
func (ctrl *HotelController) GetHotel(c *gin.Context) {
	id := c.Param("id")

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	hotel, err := services.GetHotel(objectID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

// Obtener todos los hoteles
func (ctrl *HotelController) GetHotels(c *gin.Context) {
	hotels, err := services.GetHotels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hotels"})
		return
	}

	c.JSON(http.StatusOK, hotels)
}

// Actualizar un hotel
func (ctrl *HotelController) UpdateHotel(c *gin.Context) {
	id := c.Param("id")
	var hotelDto models.Hotel

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	if err := c.ShouldBindJSON(&hotelDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel, err := services.UpdateHotel(objectID, hotelDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hotel"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

// Eliminar un hotel
func (ctrl *HotelController) DeleteHotel(c *gin.Context) {
	id := c.Param("id")

	// Convertir el ID de string a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	err = services.DeleteHotel(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
