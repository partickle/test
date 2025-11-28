package pr

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const tablePullRequest = "pull_requests"
const tablePullRequestReviewers = "pull_requests_reviewers"

type Repository struct {
	pool         *pgxpool.Pool
	queryBuilder squirrel.StatementBuilderType
}

func NewPrRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool:         pool,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
