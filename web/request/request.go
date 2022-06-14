package request

type UserCreateRequest struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type SendVerifyRequest struct {
	Email string `json:"email" form:"email"`
}

type ProfileUpdateRequest struct {
	FirstName   string `json:"first_name" form:"firstname"`
	LastName    string `json:"last_name" form:"lastname"`
	Email       string `json:"email" form:"email"`
	Address     string `json:"address" form:"address"`
	PostalCode  string `json:"postal_code" form:"postalcode"`
	City        string `json:"city" form:"city"`
	State       string `json:"state" form:"state"`
	Country     string `json:"country" form:"country"`
	PhoneNumber string `json:"phone_number" form:"phonenumber"`
}

type InvoiceCreateRequest struct {
	InvoiceNumber   string                        `json:"invoice_number"`
	Description     string                        `json:"description"`
	TypePayment     string                        `json:"type_payment"`
	Date            string                        `json:"date"`
	DateDue         string                        `json:"date_due"`
	Total           float64                       `json:"total"`
	SubTotal        float64                       `json:"sub_total"`
	LogoURL         string                        `json:"logo_url"`
	Items           []ItemCreateRequest           `json:"items"`
	AdditionalCosts []AdditionalCostCreateRequest `json:"additional_costs"`
	Client          ClientCreateRequest           `json:"client"`
}

type ItemCreateRequest struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type AdditionalCostCreateRequest struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Total float64 `json:"total"`
}

type ClientCreateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone_number"`
}
