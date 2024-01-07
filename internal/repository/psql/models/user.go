package models

import "time"

type User struct {
	Id             string    `dbq:"id"`
	Name           string    `dbq:"name"`
	Email          string    `dbq:"email"`
	Password       string    `dbq:"password"`
	Position       string    `dbq:"position"`
	PhoneNumber    string    `dbq:"phone_number"`
	CompanyName    string    `dbq:"company_name"`
	CompanyAddress string    `dbq:"company_address"`
	Information    string    `dbq:"information"`
	CreatedAt      time.Time `dbq:"created_at"`
	UpdatedAt      time.Time `dbq:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func TableUser() []string {
	return []string{
		"id",
		"name",
		"email",
		"password",
		"position",
		"phone_number",
		"company_name",
		"company_address",
		"information",
		"created_at",
		"updated_at",
	}
}
