package main

import (
	"hotel-api/initializers"
	"hotel-api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	// Cargar variables de entorno
	// initializers.LoadEnvVariables() // Si usas .env, puedes habilitar esta línea

	// Conectar a la base de datos
	initializers.ConnectMongo()

	// Conectar a RabbitMQ
	err := initializers.ConnectRabbitMQ() // Llamamos a la función para inicializar RabbitMQ
	if err != nil {
		panic("Failed to connect to RabbitMQ") // En caso de error, paramos la aplicación
	}
}

func main() {
	r := gin.Default()

	// Configuración CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"}, // Cambia esto por el origen correcto de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Llamar al archivo de rutas para registrar las rutas de los hoteles y amenidades
	routes.SetupHotelRoutes(r)
	routes.SetupAmenityRoutes(r)

	// Iniciar el servidor
	r.Run(":8080")
}
