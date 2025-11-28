package team

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/team"
)

type Service struct {
	repo teamRepository
}

func NewTeamService(repo teamRepository) *Service {
	return &Service{repo}
}

func (s *Service) Add(ctx context.Context, team team.Team) (*team.Team, error) {
	return s.repo.Add(ctx, team)
}

func (s *Service) Get(ctx context.Context, teamName string) (*team.Team, error) {
	return s.repo.Get(ctx, teamName)
}
