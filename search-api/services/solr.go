package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"search-api/dtos"
)

const solrURL = "http://localhost:8983/solr/hotels_core"

// AddHotel se encarga de agregar un hotel a Solr
func AddHotel(hotelDTO dtos.HotelDTO) error {
	// Preparar el hotel en formato JSON
	doc := map[string]interface{}{
		"name":      hotelDTO.Name,
		"location":  hotelDTO.Location,
		"amenities": hotelDTO.Amenities,
		"rating":    hotelDTO.Rating,
	}

	// Convertir a JSON
	data, err := json.Marshal([]map[string]interface{}{doc})
	if err != nil {
		return fmt.Errorf("error marshaling hotel: %v", err)
	}

	// Hacer la solicitud a Solr para indexar el documento
	req, err := http.NewRequest("POST", solrURL+"/update?commit=true", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error creating request to Solr: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request to Solr: %v", err)
	}
	defer resp.Body.Close()

	// Verificar si la respuesta es exitosa
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error indexing hotel, status code: %d", resp.StatusCode)
	}

	return nil
}
