package solr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura de ejemplo para los documentos de hoteles
type Hotel struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Enviar el hotel a Solr
func SendToSolr(message []byte) error {
	// Deserializar el mensaje recibido
	var hotel Hotel
	err := json.Unmarshal(message, &hotel)
	if err != nil {
		return fmt.Errorf("Error unmarshaling message: %v", err)
	}

	// Convertir el hotel a JSON para Solr
	jsonData, err := json.Marshal(map[string]interface{}{
		"add": map[string]interface{}{
			"doc": hotel,
		},
	})
	if err != nil {
		return fmt.Errorf("Error marshaling data to Solr: %v", err)
	}

	// Enviar el documento a Solr
	solrURL := "http://localhost:8983/solr/hotel_core/update?commit=true"
	req, err := http.NewRequest("POST", solrURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request to Solr: %v", err)
	}
	defer resp.Body.Close()

	// Verificar la respuesta de Solr
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Solr returned an error: %v", resp.Status)
	}

	log.Println("Document successfully indexed in Solr.")
	return nil
}
