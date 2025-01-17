package service

import "context"

type UserRepository interface {
	Register(ctx context.Context)
	GetToken(ctx context.Context, username, pass string) (string, error)
}

type UserService struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context) {}
func (s *UserService) GetToken(ctx context.Context, username, pass string) (string, error) {
	return s.Repo.GetToken(ctx, username, pass)
}
