package repository

import (
	"context"
	"github.com/IlyaZayats/shopus/internal/entity"
	"github.com/IlyaZayats/shopus/internal/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresListRepository struct {
	db *pgxpool.Pool
}

func NewPostgresListRepository(db *pgxpool.Pool) (interfaces.ListRepository, error) {
	return &PostgresListRepository{
		db: db,
	}, nil
}

func (r *PostgresListRepository) GetLists() ([]entity.List, error) {
	var lists []entity.List
	q := "select row_to_json(x) from (select * from Lists) as x"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return lists, err
	}
	for rows.Next() {
		var list entity.List
		if err := rows.Scan(&list); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil

}

func (r *PostgresListRepository) InsertList(list entity.List) error {
	q := "INSERT INTO Lists (dealer_id, name, price, amount, created_at, info, carrier, contact_person, note) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	if _, err := r.db.Exec(context.Background(), q, list.DealerId, list.Name, list.Price, list.Amount, list.CreatedAt, list.Info, list.Carrier, list.ContactPerson, list.Note); err != nil {
		return err
	}
	return nil
}

func (r *PostgresListRepository) UpdateList(list entity.List) error {
	q := "UPDATE Lists SET (name, price, amount, created_at, info, carrier, contact_person, note) = ($1, $2, $3, $4, $5, $6, $7, $8) WHERE id=$9"
	if _, err := r.db.Exec(context.Background(), q, list.Name, list.Price, list.Amount, list.CreatedAt, list.Info, list.Carrier, list.ContactPerson, list.Note, list.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresListRepository) DeleteList(id int) error {
	q := "DELETE FROM Lists WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}
