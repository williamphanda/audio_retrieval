package userphrase

import (
	entityphrase "audioretrieval/internal/entity/phrase"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// GetByUserIDAndPhraseID get data from database from user id and phrase id
func (r *Repo) GetByID(phraseID int64) (entityphrase.Phrase, error) {
	var phrase entityphrase.Phrase

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.
		Select("id").
		From("audio_retrieval_trx_phrases").
		Where(squirrel.Eq{"id": phraseID}).
		ToSql()
	if err != nil {
		return phrase, fmt.Errorf("failed to create get phrase query: %v", err)
	}

	rows, err := r.dbInterface.Query(query, args...)
	if err != nil {
		return phrase, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return phrase, fmt.Errorf("phrase ID %d not found", phraseID)
	}

	err = rows.Scan(&phrase.ID)
	if err != nil {
		return phrase, fmt.Errorf("failed to scan row into struct: %v", err)
	}

	return phrase, nil
}
