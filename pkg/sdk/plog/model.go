package plog

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ProjectId  *string    `json:"project_id,omitempty"`
	Identifier *string    `json:"id"`
	Type       *string    `json:"type"`
	Message    *string    `json:"message"`
	CreatedAt  *time.Time `json:"created_at"`
}

type LogEvent struct {
	Log
	EventType string `json:"event_type"`
}

func (l Log) IsValid() bool {
	if l.Identifier == nil || l.Type == nil || l.Message == nil {
		return false
	}

	if *l.Message == "" {
		return false
	}

	if err := uuid.Validate(*l.Identifier); err != nil {
		return false
	}

	if !slices.Contains(LogTypes[:], *l.Type) {
		return false
	}

	return true
}

func (l LogEvent) IsValid() bool {
	if l.ProjectId == nil {
		return false
	}

	if err := uuid.Validate(*l.ProjectId); err != nil {
		return false
	}

	if ok := l.Log.IsValid(); !ok {
		return false
	}

	if !slices.Contains(LogEvents[:], l.EventType) {
		return false
	}

	return true
}
