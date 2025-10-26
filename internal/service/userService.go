package service

import (
	"context"
	"time"

	"github.com/Anagh3/go-backend/internal/repository"
)

type UserWithAge struct {
	ID   uint64    `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age"`
}

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser(ctx context.Context, name string, dob time.Time) error {
	return s.repo.CreateUser(ctx, name, dob)
}

func (s *UserService) GetUser(ctx context.Context, id uint64) (UserWithAge, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return UserWithAge{}, err
	}
	return UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob,
		Age:  calculateAge(user.Dob),
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]UserWithAge, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var result []UserWithAge
	for _, u := range users {
		result = append(result, UserWithAge{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob,
			Age:  calculateAge(u.Dob),
		})
	}
	return result, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint64, name string, dob time.Time) error {
	return s.repo.UpdateUser(ctx, id, name, dob)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint64) error {
	return s.repo.DeleteUser(ctx, id)
}



func calculateAge(dob time.Time) int {
	now := time.Now()

	// Extract only year, month, day (ignore time and timezone)
	y1, m1, d1 := now.Date()
	y2, m2, d2 := dob.Date()

	age := y1 - y2
	// Subtract 1 if birthday hasn't occurred yet this year
	if m1 < m2 || (m1 == m2 && d1 < d2) {
		age--
	}
	return age
}
