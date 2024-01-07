package request

type RegisterUser struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Position       string `json:"position"`
	PhoneNumber    string `json:"phone_number"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	Information    string `json:"information"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
