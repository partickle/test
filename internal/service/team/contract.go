package team

import (
	"context"

	"github.com/partickle/avito-pr-review-service/internal/model/team"
)

type teamRepository interface {
	Add(ctx context.Context, team team.Team) (*team.Team, error)
	Get(ctx context.Context, teamName string) (*team.Team, error)
}
