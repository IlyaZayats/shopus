package interfaces

import "github.com/IlyaZayats/shopus/internal/entity"

type ListRepository interface {
	GetLists() ([]entity.List, error)
	UpdateList(List entity.List) error
	InsertList(List entity.List) error
	DeleteList(id int) error
}
