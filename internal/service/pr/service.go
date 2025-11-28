package pr

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/pr"
)

type Service struct {
	repo prRepository
}

func NewPrService(repo prRepository) *Service {
	return &Service{repo}
}

func (s *Service) Create(ctx context.Context, pr pr.PullRequest) (*pr.PullRequest, error) {
	return s.repo.Create(ctx, pr)
}
func (s *Service) Merge(ctx context.Context, prID string) (*pr.PullRequest, error) {
	return s.repo.Merge(ctx, prID)
}
func (s *Service) Reassign(ctx context.Context, prID, oldUserID string) (string, *pr.PullRequest, error) {
	return s.repo.Reassign(ctx, prID, oldUserID)
}
