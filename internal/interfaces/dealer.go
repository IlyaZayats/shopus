package interfaces

import "github.com/IlyaZayats/shopus/internal/entity"

type DealerRepository interface {
	GetDealers() ([]entity.Dealer, error)
	UpdateDealer(Dealer entity.Dealer) error
	InsertDealer(Dealer entity.Dealer) error
	DeleteDealer(id int) error
}
