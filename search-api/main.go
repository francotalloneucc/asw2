package main

import (
	"fmt"
	"log"
	"search-api/consumer"
)

func main() {
	// Iniciar el consumidor para escuchar la cola de RabbitMQ
	fmt.Println("Iniciando el consumidor de RabbitMQ...")

	// Llamar al método Consume para que comience a recibir y procesar los mensajes
	err := consumer.Consume()
	if err != nil {
		log.Fatalf("Error al consumir los mensajes: %v", err)
	}
}
