package db

import (
	"context"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/kemper0110/bus-schedule/internal/models"
	_ "github.com/lib/pq"
)

type Store struct {
	conn *sqlx.DB
}

func NewDB(connURI string) (*Store, error) {
	conn, err := sqlx.Connect("postgres", connURI)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	if err := startMigration(conn); err != nil {
		return nil, err
	}
	return &Store{
		conn: conn,
	}, nil
}

func startMigration(conn *sqlx.DB) error {
	driver, err := postgres.WithInstance(conn.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func (s *Store) GetCitiesList(ctx context.Context) ([]models.City, error) {
	var c []models.City
	if err := s.conn.SelectContext(ctx, &c, `
	SELECT *
	FROM city
	`); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Store) GetCarriersList(ctx context.Context) ([]models.Carriers, error) {
	var c []models.Carriers
	if err := s.conn.SelectContext(ctx, &c, `
	SELECT *
	FROM carriers
	`); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Store) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	var us models.User
	if err := s.conn.GetContext(ctx, &us, `
	INSERT INTO users(login, password)
	VALUES ($1, $2)
	RETURNING id, role;
	`, user.Login, user.Password); err != nil {
		return models.User{}, err
	}
	us.Login = user.Login
	return us, nil
}

func (s *Store) GetUserIDByLogin(ctx context.Context, user models.User) (models.User, error) {
	var us models.User
	if err := s.conn.GetContext(ctx, &us, `
	SELECT *
	FROM users
	WHERE login = $1
	`, user.Login); err != nil {
		return models.User{}, err
	}
	return us, nil
}

func (s *Store) GetRoutes(ctx context.Context, route models.SerchParams) ([]models.MidRoute, error) {
	var rt []models.SearchRoute
	if err := s.conn.SelectContext(ctx, &rt, `
	SELECT mid_route.id, main_id, from_city, to_city, from_time, to_time, price, rating , c.name
	FROM mid_route
	INNER JOIN main_route mr on mr.id = mid_route.main_id
	INNER JOIN carriers c on mr.carrier_id = c.id
	WHERE from_city =  AND to_city = $2 AND from_time = $3
	`, route); err != nil {
		return nil, err
	}
	return rt, nil
}
