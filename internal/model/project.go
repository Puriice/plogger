package model

import "time"

type Project struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
