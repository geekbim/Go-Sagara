package models

import (
	"time"
)

type Region struct {
	Id        int       `dbq:"id"`
	Name      string    `dbq:"name"`
	CreatedAt time.Time `dbq:"created_at"`
	UpdatedAt time.Time `dbq:"updated_at"`
}

type RegionList struct {
	Id        int       `dbq:"id"`
	Name      string    `dbq:"name"`
	CreatedAt time.Time `dbq:"created_at"`
	UpdatedAt time.Time `dbq:"updated_at"`
}

func (Region) TableName() string {
	return "indonesia_provinces"
}

func TableRegion() []string {
	return []string{
		"id",
		"name",
		"created_at",
		"updated_at",
	}
}
