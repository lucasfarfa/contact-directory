package contacts

type Contact struct {
	ID      int      `json:"contact_id"`
	Name    string   `json:"first_name"`
	Surname string   `json:"last_name"`
	Company Company  `json:"company"`
	Email   string   `json:"email"`
	Phones  []string `json:"phones"`
}

type Company struct {
	ID     int    `json:"company_id"`
	Name   string `json:"name"`
	Addres string `json:"address"`
}
