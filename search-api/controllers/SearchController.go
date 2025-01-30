package controllers

import (
	"log"
	"net/http"
	"search-api/dtos"
	"search-api/services"
	"github.com/gin-gonic/gin"
)

// IndexHotel maneja la indexación de un nuevo hotel en Solr
func IndexHotel(c *gin.Context) {
	var hotelDTO dtos.HotelDTO

	// Bindear el cuerpo JSON a la estructura de HotelDTO
	if err := c.ShouldBindJSON(&hotelDTO); err != nil {
		log.Println("❌ Error al bindear JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Llamar al servicio para indexar el hotel en Solr
	if err := services.AddHotel(hotelDTO); err != nil {
		log.Println("❌ Error al indexar hotel:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to index hotel"})
		return
	}

	// Responder al cliente indicando éxito
	c.JSON(http.StatusOK, gin.H{"message": "Hotel indexed successfully!"})
}
