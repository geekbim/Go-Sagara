package request

import "time"

type CreateOrder struct {
	PicName          string    `json:"pic_name"`
	DateStart        time.Time `json:"date_start"`
	DateEnd          time.Time `json:"date_end"`
	Payment          string    `json:"payment"`
	CapacityEmployee int       `json:"capacity_employee"`
	Requirement      string    `json:"requirement"`
	Other            string    `json:"other"`
}
