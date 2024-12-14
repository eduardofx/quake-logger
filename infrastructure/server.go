package infrastructure

import (
	"log"
	"net/http"
	"quake-logger/application"
	"quake-logger/interfaces"
)

func StartServer() {
	service := application.NewLogParserService()
	if err := service.ParseLogFile("shared/data/qgames.log"); err != nil {
		log.Fatalf("Failed to parse log file: %v", err)
	}

	handler := interfaces.NewLogHandler(service)

	http.HandleFunc("/reports", handler.HandleReports)
	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
