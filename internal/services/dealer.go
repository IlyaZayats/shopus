package services

import (
	"encoding/json"
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
)

type DealerService struct {
	repo interfaces.DealerRepository
}

func NewDealerService(repo interfaces.DealerRepository) (*DealerService, error) {
	return &DealerService{
		repo: repo,
	}, nil
}

func (s *DealerService) GetDealers() ([]map[string]interface{}, error) {
	dealers, err := s.repo.GetDealers()
	if err != nil {
		return nil, err
	}
	var dealersSlice []map[string]interface{}
	for _, item := range dealers {
		tmpByte, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		var tmpMap map[string]interface{}
		if err := json.Unmarshal(tmpByte, &tmpMap); err != nil {
			return nil, err
		}
		dealersSlice = append(dealersSlice, tmpMap)
	}
	return dealersSlice, nil
}

func (s *DealerService) InsertDealer(networkId int, name string) error {
	return s.repo.InsertDealer(entity.Dealer{Id: 0, NetworkId: networkId, Name: name})
}

func (s *DealerService) UpdateDealer(id int, name string) error {
	return s.repo.UpdateDealer(entity.Dealer{Id: id, NetworkId: 0, Name: name})
}

func (s *DealerService) DeleteDealer(id int) error {
	return s.repo.DeleteDealer(id)
}
