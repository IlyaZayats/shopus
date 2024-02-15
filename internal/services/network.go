package services

import (
	"encoding/json"
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
)

type NetworkService struct {
	repo interfaces.NetworkRepository
}

func NewNetworkService(repo interfaces.NetworkRepository) (*NetworkService, error) {
	return &NetworkService{
		repo: repo,
	}, nil
}

func (s *NetworkService) GetNetworks() ([]map[string]interface{}, error) {
	networks, err := s.repo.GetNetworks()
	if err != nil {
		return nil, err
	}
	var networksSlice []map[string]interface{}
	for _, item := range networks {
		tmpByte, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		var tmpMap map[string]interface{}
		if err := json.Unmarshal(tmpByte, &tmpMap); err != nil {
			return nil, err
		}
		networksSlice = append(networksSlice, tmpMap)
	}
	return networksSlice, nil
}

func (s *NetworkService) InsertNetwork(name string) error {
	return s.repo.InsertNetwork(entity.Network{Id: 0, Name: name})
}

func (s *NetworkService) UpdateNetwork(id int, name string) error {
	return s.repo.UpdateNetwork(entity.Network{Id: id, Name: name})
}

func (s *NetworkService) DeleteNetwork(id int) error {
	return s.repo.DeleteNetwork(id)
}
