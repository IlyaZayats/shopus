package services

import (
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
	"github.com/IlyaZayats/shopus/internal/requests"
	"strconv"
)

type ListService struct {
	repo interfaces.ListRepository
}

func NewListService(repo interfaces.ListRepository) (*ListService, error) {
	return &ListService{
		repo: repo,
	}, nil
}

func (s *ListService) GetLists() ([]map[string]string, error) {
	lists, err := s.repo.GetLists()
	if err != nil {
		return nil, err
	}
	var listsSlice []map[string]string
	for _, item := range lists {
		tmpMap := map[string]string{
			"id":             strconv.Itoa(item.Id),
			"dealer_id":      strconv.Itoa(item.DealerId),
			"name":           item.Name,
			"price":          strconv.Itoa(item.Price),
			"amount":         strconv.Itoa(item.Amount),
			"created_at":     item.CreatedAt,
			"info":           item.Info,
			"carrier":        item.Carrier,
			"contact_person": item.ContactPerson,
			"note":           item.Note,
		}
		listsSlice = append(listsSlice, tmpMap)
	}
	return listsSlice, nil
}

func (s *ListService) InsertList(request requests.InsertListRequest) error {
	return s.repo.InsertList(entity.List(request))
}

func (s *ListService) UpdateList(request requests.UpdateListRequest) error {
	return s.repo.UpdateList(entity.List(request))
}

func (s *ListService) DeleteList(id int) error {
	return s.repo.DeleteList(id)
}
