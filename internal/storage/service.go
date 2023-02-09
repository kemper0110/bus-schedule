package db

import (
	"context"

	"github.com/kemper0110/bus-schedule/models"
)

type Service struct {
	db db
}

func (s *Service) CitiesList(ctx context.Context) ([]models.City, error) {
	list, err := s.db.getCitiesList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) CarriersList(ctx context.Context) ([]models.Carriers, error) {
	list, err := s.db.getCarriersList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) SignUp(ctx context.Context, user models.User)
