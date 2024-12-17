package user

import (
	"fmt"

	entityuser "audioretrieval/internal/entity/user"

	"github.com/Masterminds/squirrel"
)

// GetByID get data from database for a selected user ID
func (r *Repo) GetByID(userID int64) (entityuser.User, error) {
	var user entityuser.User

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.
		Select("id").
		From("audio_retrieval_trx_users").
		Where(squirrel.Eq{"id": userID}).
		ToSql()
	if err != nil {
		return user, fmt.Errorf("failed to create insert user query: %v", err)
	}

	rows, err := r.dbInterface.Query(query, args...)
	if err != nil {
		return user, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return user, fmt.Errorf("user ID %d not found", userID)
	}

	err = rows.Scan(&user.ID)
	if err != nil {
		return user, fmt.Errorf("failed to scan row into struct: %v", err)
	}

	return user, nil
}
