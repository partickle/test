package user

import (
	"github.com/partickle/avito-pr-review-service/internal/model/pr"
	"github.com/partickle/avito-pr-review-service/internal/model/user"
)

func ModelToResponse(user user.User) SetIsActiveResponse {
	return SetIsActiveResponse{
		UserID:   user.UserID,
		Username: user.Username,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}

func ModelToResponse(user user.User, pr pr.PullRequest) GetReviewResponse {
	return GetReviewResponse{}
}
