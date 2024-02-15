package requests

//type GetListsRequest struct {
//	Id int `json:"id" binding:"required"`
//}

type InsertListRequest struct {
	Id            int    `json:"id"`
	DealerId      int    `json:"dealer_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	CreatedAt     string `json:"created_at" binding:"required"`
	Info          string `json:"info" binding:"required"`
	Carrier       string `json:"carrier" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Note          string `json:"Note" binding:"required"`
}

type UpdateListRequest struct {
	Id            int    `json:"id" binding:"required"`
	DealerId      int    `json:"dealer_id"`
	Name          string `json:"name" binding:"required"`
	Price         int    `json:"price" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
	CreatedAt     string `json:"created_at" binding:"required"`
	Info          string `json:"info" binding:"required"`
	Carrier       string `json:"carrier" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Note          string `json:"Note" binding:"required"`
}

type DeleteListRequest struct {
	Id int `json:"id" binding:"required"`
}
