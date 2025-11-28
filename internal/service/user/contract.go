package user

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/user"
)

type userRepository interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*user.User, error)
	GetReview(ctx context.Context, userID string) (*models.UserReviewResponse, error)
}
