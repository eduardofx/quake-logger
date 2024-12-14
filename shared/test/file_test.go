package test

import (
	"quake-logger/application"
	"testing"
)

func TestParseLogFile_NonExistentFile(t *testing.T) {
	service := application.NewLogParserService()
	err := service.ParseLogFile("null.log")
	if err == nil {
		t.Fatalf("expected an error for non-existent file, got nil")
	}
}
