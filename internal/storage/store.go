package db

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kemper0110/bus-schedule/internal/models"
	_ "github.com/lib/pq"
)

type db struct {
	conn *sqlx.DB
}

func NewDB(connURI string) (*db, error) {
	conn, err := sqlx.Connect("postgres", connURI)
	if err != nil {
		return nil, err
	}
	startMigration(conn)
	return &db{
		conn: conn,
	}, nil
}

func startMigration(conn *sqlx.DB) error {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	m.Up()
	return nil
}

func (d *db) getCitiesList(ctx context.Context) ([]models.City, error) {
	var c []models.City
	if err := d.conn.SelectContext(ctx, c, `
	SELECT *
	FROM city
	`); err != nil {
		return nil, err
	}
	return c, nil
}

func (d *db) getCarriersList(ctx context.Context) ([]models.Carriers, error) {
	var c []models.Carriers
	if err := d.conn.SelectContext(ctx, c, `
	SELECT *
	FROM carriers
	`); err != nil {
		return nil, err
	}
	return c, nil
}
