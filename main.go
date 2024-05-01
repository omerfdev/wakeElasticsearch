package main

import (
	"fmt"
	"log"
	"net/http"	
	"os/exec"
	"time"
)

func main() {
	// HTTP endpoint'i tanımla
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		// Elasticsearch ve Kibana'yı başlat
		err := startElasticsearchAndKibana()
		if err != nil {
			http.Error(w, "Failed to start Elasticsearch and Kibana", http.StatusInternalServerError)
			return
		}

		// Başarı mesajını gönder
		w.Write([]byte("Elasticsearch and Kibana started successfully"))
	})

	// HTTP sunucusunu başlat
	fmt.Println("HTTP server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startElasticsearchAndKibana() error {
	// Elasticsearch ve Kibana'yı başlatmak için komutları tanımla
	elasticsearchCmd := exec.Command("elasticsearch")
	kibanaCmd := exec.Command("kibana")

	// Komutları arka planda çalıştır
	err1 := elasticsearchCmd.Start()
	err2 := kibanaCmd.Start()

	// Hata kontrolü yap
	if err1 != nil || err2 != nil {
		return fmt.Errorf("Failed to start Elasticsearch or Kibana: %v, %v", err1, err2)
	}

	// Elasticsearch ve Kibana'nın başlaması için biraz bekleyelim
	time.Sleep(5 * time.Second)

	return nil
}
