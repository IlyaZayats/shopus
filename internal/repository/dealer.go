package repository

import (
	"context"
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDealerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresDealerRepository(db *pgxpool.Pool) (interfaces.DealerRepository, error) {
	return &PostgresDealerRepository{
		db: db,
	}, nil
}

func (r *PostgresDealerRepository) GetDealers() ([]entity.Dealer, error) {
	var dealers []entity.Dealer
	q := "select row_to_json(x) from (select * from Dealers) as x"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return dealers, err
	}
	defer rows.Close()
	for rows.Next() {
		var dealer entity.Dealer
		if err := rows.Scan(&dealer); err != nil {
			return nil, err
		}
		dealers = append(dealers, dealer)
	}
	return dealers, nil

}

func (r *PostgresDealerRepository) InsertDealer(dealer entity.Dealer) error {
	q := "INSERT INTO Dealers (network_id, name) VALUES ($1, $2)"
	if _, err := r.db.Exec(context.Background(), q, dealer.NetworkId, dealer.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresDealerRepository) UpdateDealer(dealer entity.Dealer) error {
	q := "UPDATE Dealers SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, dealer.Name, dealer.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresDealerRepository) DeleteDealer(id int) error {
	q := "DELETE FROM Dealers WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}
