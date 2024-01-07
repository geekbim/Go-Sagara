package entity

import (
	"time"
)

type Region struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
