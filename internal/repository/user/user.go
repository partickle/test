package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/partickle/avito-pr-review-service/internal/model/user"
)

const tableUsers = "users"

type Repository struct {
	pool         *pgxpool.Pool
	queryBuilder squirrel.StatementBuilderType
}

func NewUserRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool:         pool,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *Repository) SetIsActive(ctx context.Context, userID string, isActive bool) (*user.User, error) {
	query := r.queryBuilder.
		Update(tableUsers).
		Set("is_active", isActive).
		Where(squirrel.Eq{"id": userID}).
		Suffix("RETURNING user_id, username, team_name, is_active")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var u user.User
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&u.UserID,
		&u.Username,
		&u.TeamName,
		&u.IsActive,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return &u, nil
}
