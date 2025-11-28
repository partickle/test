package team

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const tableTeam = "teams"
const tableTeamMember = "team_members"

type Repository struct {
	pool         *pgxpool.Pool
	queryBuilder squirrel.StatementBuilderType
}

func NewTeamRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool:         pool,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
