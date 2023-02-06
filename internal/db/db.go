package db

import (
	"context"

	"github.com/IB133/bus-tickets/internal/models"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn *sqlx.DB
}

func NewDB(connURI string) (*DB, error) {
	conn, err := sqlx.Connect("postgres", connURI)
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

func (d *DB) getCitiesList(ctx context.Context) ([]models.City, error) {
	var c []models.City
	if err := d.conn.SelectContext(ctx, c, `
	SELECT *
	FROM city
	`); err != nil {
		return nil, err
	}
	return c, nil
}

func (d *DB) getCarriersList(ctx context.Context) ([]models.Carriers, error) {
	var c []models.Carriers
	if err := d.conn.SelectContext(ctx, c, `
	SELECT *
	FROM carriers
	`); err != nil {
		return nil, err
	}
	return c, nil
}
