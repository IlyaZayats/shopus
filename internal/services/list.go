package services

import (
	"encoding/json"
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
	"github.com/IlyaZayats/shopus/internal/requests"
)

type ListService struct {
	repo interfaces.ListRepository
}

func NewListService(repo interfaces.ListRepository) (*ListService, error) {
	return &ListService{
		repo: repo,
	}, nil
}

func (s *ListService) GetLists() ([]map[string]interface{}, error) {
	lists, err := s.repo.GetLists()
	if err != nil {
		return nil, err
	}
	var listsSlice []map[string]interface{}
	for _, item := range lists {
		tmpByte, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		var tmpMap map[string]interface{}
		if err := json.Unmarshal(tmpByte, &tmpMap); err != nil {
			return nil, err
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
