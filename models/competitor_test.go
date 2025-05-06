package models

import (
	"testing"
	"time"
)

func TestLapTiming(t *testing.T) {
	comp := Competitor{ID: 1}

	start := time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 1, 9, 5, 0, 0, time.UTC)

	comp.LapStartTimes = append(comp.LapStartTimes, start)
	comp.LapTimes = append(comp.LapTimes, end.Sub(start))

	if len(comp.LapTimes) != 1 {
		t.Fatalf("Expected 1 lap time, got %d", len(comp.LapTimes))
	}
	expected := 5 * time.Minute
	if comp.LapTimes[0] != expected {
		t.Errorf("Expected lap time %v, got %v", expected, comp.LapTimes[0])
	}
}

func TestShotTracking(t *testing.T) {
	comp := Competitor{ID: 1}
	comp.ShotsFired = 5
	comp.Hits = 3

	if comp.ShotsFired != 5 || comp.Hits != 3 {
		t.Errorf("Expected 5 shots, 3 hits; got %d shots, %d hits", comp.ShotsFired, comp.Hits)
	}
}
