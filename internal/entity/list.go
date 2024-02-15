package entity

type List struct {
	Id            int    `json:"id"`
	DealerId      int    `json:"dealer_id"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Amount        int    `json:"amount"`
	CreatedAt     string `json:"created_at"`
	Info          string `json:"info"`
	Carrier       string `json:"carrier"`
	ContactPerson string `json:"contact_person"`
	Note          string `json:"note"`
}
