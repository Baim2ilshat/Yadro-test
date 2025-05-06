package config

import (
	"testing"
	"time"
)

func TestParseLine(t *testing.T) {
	line := "[09:30:01.005] 4 1"
	event, err := parseLine(line)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedTime, _ := time.Parse("15:04:05.000", "09:30:01.005")

	if event.Time != expectedTime {
		t.Errorf("Expected time %v, got %v", expectedTime, event.Time)
	}
	if event.EventID != 4 {
		t.Errorf("Expected EventID 4, got %d", event.EventID)
	}
	if event.CompetitorID != 1 {
		t.Errorf("Expected CompetitorID 1, got %d", event.CompetitorID)
	}
	if event.Extra != "" {
		t.Errorf("Expected Extra to be empty, got '%s'", event.Extra)
	}
}
