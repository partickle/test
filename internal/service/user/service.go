package user

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/user"
)

type Service struct {
	repo userRepository
}

func NewUserService(repo userRepository) *Service {
	return &Service{repo}
}

func (s *Service) SetIsActive(ctx context.Context, userID string, isActive bool) (*user.User, error) {
	return s.repo.SetIsActive(ctx, userID, isActive)
}

func GetReview(ctx context.Context, userID string) (*models.UserReviewResponse, error) {

}
