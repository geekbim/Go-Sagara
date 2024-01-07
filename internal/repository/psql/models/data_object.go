package models

import "time"

type DataObject struct {
	Id        string    `dbq:"id"`
	Path      string    `dbq:"path"`
	CreatedAt time.Time `dbq:"created_at"`
	UpdatedAt time.Time `dbq:"updated_at"`
}

func (DataObject) TableName() string {
	return "data_objects"
}

func TableDataObject() []string {
	return []string{
		"id",
		"path",
		"created_at",
		"updated_at",
	}
}
