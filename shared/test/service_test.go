package test

import (
	"os"
	"quake-logger/application"
	"testing"
)

func TestLogParserService_ParseLogFile(t *testing.T) {
	// Mock file content
	mockLogContent := `
		0:00 InitGame: ...
		0:05 Kill: 1022 2 22: <world> killed Player1 by MOD_TRIGGER_HURT
		0:10 Kill: 2 3 7: Player2 killed Player3 by MOD_ROCKET
		0:15 Kill: 3 2 6: Player3 killed Player2 by MOD_SHOTGUN
		0:20 InitGame: ...
		0:25 Kill: 1022 1 22: <world> killed Player2 by MOD_TRIGGER_HURT
	`

	// Create a temporary file with the mock content
	tmpFile := createTempFile(t, mockLogContent)
	defer tmpFile.Close()

	// Instantiate the service
	service := application.NewLogParserService()

	// Parse the log file
	err := service.ParseLogFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to parse log file: %v", err)
	}

	// Verify the parsed data
	if len(service.Matches) != 2 {
		t.Errorf("Expected 2 matches, got %d", len(service.Matches))
	}

	// Check the details of the first match
	match1 := service.Matches["game_1"]
	if match1 == nil {
		t.Fatalf("Match game_1 not found")
	}
	if match1.TotalKills != 3 {
		t.Errorf("Expected 3 kills in game_1, got %d", match1.TotalKills)
	}
	if len(match1.Players) != 3 {
		t.Errorf("Expected 3 players in game_1, got %d", len(match1.Players))
	}
	if match1.Kills["Player2"] != 1 {
		t.Errorf("Expected Player2 to have 1 kill in game_1, got %d", match1.Kills["Player2"])
	}
	if match1.KillsByMeans["MOD_ROCKET"] != 1 {
		t.Errorf("Expected 1 kill by MOD_ROCKET in game_1, got %d", match1.KillsByMeans["MOD_ROCKET"])
	}
}

func createTempFile(t *testing.T, content string) *os.File {
	tmpFile, err := os.CreateTemp("", "mock-log-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	return tmpFile
}
