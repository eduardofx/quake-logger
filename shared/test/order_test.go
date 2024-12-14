package test

import (
	"quake-logger/application"
	"quake-logger/domain"
	"testing"
)

func TestGenerateReports_Order(t *testing.T) {
	service := application.NewLogParserService()
	service.Matches["game_2"] = &domain.Match{GameID: "game_2"}
	service.Matches["game_1"] = &domain.Match{GameID: "game_1"}
	service.Matches["game_3"] = &domain.Match{GameID: "game_3"}

	reports := service.GenerateReports()
	expectedOrder := []string{"game_1", "game_2", "game_3"}
	for i, report := range reports {
		if report.GameID != expectedOrder[i] {
			t.Errorf("expected game ID %s, got %s", expectedOrder[i], report.GameID)
		}
	}
}
