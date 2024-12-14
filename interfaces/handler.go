package interfaces

import (
	"encoding/json"
	"net/http"
	"quake-logger/application"
)

type GameReport struct {
	GameID string         `json:"game_id"`
	Data   map[string]int `json:"data"`
	Number int
}

type LogHandler struct {
	Service *application.LogParserService
}

func NewLogHandler(service *application.LogParserService) *LogHandler {
	return &LogHandler{Service: service}
}

func (lh *LogHandler) HandleReports(w http.ResponseWriter, r *http.Request) {
	reports := lh.Service.GenerateReports()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}
