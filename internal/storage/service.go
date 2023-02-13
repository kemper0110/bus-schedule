package db

import (
	"context"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/kemper0110/bus-schedule/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type WithClaims struct {
	jwt.StandardClaims
	ID    int    `json:"id"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

func (s *Store) CitiesList(ctx context.Context) ([]models.City, error) {
	list, err := s.GetCitiesList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Store) CarrierList(ctx context.Context) ([]models.Carriers, error) {
	list, err := s.GetCarriersList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Store) SignUp(ctx context.Context, reqUser models.User) (models.Token, error) {
	if err := reqUser.Validate(); err != nil {
		return models.Token{}, err
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), 10)
	if err != nil {
		return models.Token{}, err
	}

	reqUser.Password = string(hashPass)
	userDB, err := s.CreateUser(ctx, reqUser)
	if err != nil {
		return models.Token{}, err
	}

	token, err := generateJwtToken(userDB)
	if err != nil {
		return models.Token{}, err
	}

	return token, nil
}

func (s *Store) SignIn(ctx context.Context, reqUser models.User) (models.Token, error) {
	userDB, err := s.GetUserIDByLogin(ctx, reqUser)
	if err != nil {
		return models.Token{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(reqUser.Password)); err != nil {
		return models.Token{}, err
	}

	token, err := generateJwtToken(reqUser)
	if err != nil {
		return models.Token{}, err
	}

	return token, nil
}

func generateJwtToken(user models.User) (models.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, WithClaims{
		ID:    user.Id,
		Login: user.Login,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	tk, err := token.SignedString([]byte("qwesdfmksjow8457HDudf"))
	if err != nil {
		return models.Token{}, err
	}

	return models.Token{Token: tk}, nil
}
