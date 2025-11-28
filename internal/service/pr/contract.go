package pr

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/pr"
)

type prRepository interface {
	Create(ctx context.Context, pr pr.PullRequest) (*pr.PullRequest, error)
	Merge(ctx context.Context, prID string) (*pr.PullRequest, error)
	Reassign(ctx context.Context, prID, oldUserID string) (string, *pr.PullRequest, error)
}
