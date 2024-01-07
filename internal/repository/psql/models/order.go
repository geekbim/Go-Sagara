package models

import "time"

type Order struct {
	Id               string    `dbq:"id"`
	PicName          string    `dbq:"pic_name"`
	DateStart        time.Time `dbq:"date_start"`
	DateEnd          time.Time `dbq:"date_end"`
	Payment          string    `dbq:"payment"`
	CapacityEmployee int       `dbq:"capacity_employee"`
	Requirement      string    `dbq:"requirement"`
	Other            string    `dbq:"other"`
	CreatedAt        time.Time `dbq:"created_at"`
	UpdatedAt        time.Time `dbq:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

func TableOrder() []string {
	return []string{
		"id",
		"pic_name",
		"date_start",
		"date_end",
		"payment",
		"capacity_employee",
		"requirement",
		"other",
		"created_at",
		"updated_at",
	}
}
