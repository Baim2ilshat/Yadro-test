package processor

import (
	"testing"
	"time"

	"Yadro-test/config"
)

func mockEvent(timeStr string, eventID, compID int, extra string) config.Event {
	t, err := time.Parse("15:04:05.000", timeStr)
	if err != nil {
		panic(err)
	}
	return config.Event{
		Time:         t,
		EventID:      eventID,
		CompetitorID: compID,
		Extra:        extra,
	}
}

func TestProcessor_ProcessRegistrationEvent(t *testing.T) {
	cfg := &config.Config{Laps: 2, LapLen: 3651, PenaltyLen: 50}
	p := NewProcessor(cfg)

	event := mockEvent("09:05:59.867", 1, 1, "")
	p.ProcessEvent(event)

	comp := p.competitors[1]
	if comp == nil || !comp.Registered {
		t.Errorf("Expected competitor 1 to be registered")
	}
}
