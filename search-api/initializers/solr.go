package initializers

import (
	"log"
	"net/http"
	"os"
)

var SolrURL string

// InitSolr establece la conexión con Solr
func InitSolr() {
	// Obtener la URL de Solr desde las variables de entorno
	SolrURL = os.Getenv("SOLR_URL")
	if SolrURL == "" {
		log.Fatal("❌ SOLR_URL no está configurado")
	}

	// Realizar un ping para verificar la conexión
	resp, err := http.Get(SolrURL + "/admin/ping")
	if err != nil || resp.StatusCode != 200 {
		log.Fatalf("❌ Error conectando a Solr: %v", err)
	}
	defer resp.Body.Close()

	log.Println("✅ Conectado a Solr exitosamente!")
}
