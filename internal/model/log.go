package model

import (
	"time"
)

type Log struct {
	Identifier *string    `json:"id"`
	Type       *string    `json:"type"`
	Message    *string    `json:"message"`
	CreatedAt  *time.Time `json:"createdAt"`
}
