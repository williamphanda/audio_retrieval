package phrase

import (
	entityuserphrase "audioretrieval/internal/entity/userphrase"
	"fmt"

	"github.com/Masterminds/squirrel"
)

// Insert add new row of user phrase
func (r *Repo) Insert(userPhrase entityuserphrase.UserPhrase) error {
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.
		Insert("audio_retrieval_trx_user_phrases").
		Columns("user_id", "phrase_id", "file_path", "mime_type").
		Values(userPhrase.UserID, userPhrase.PhraseID, userPhrase.FilePath, userPhrase.MimeType).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to create insert phrase query: %v", err)
	}

	_, err = r.dbInterface.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute insert query and scan returned ID: %v", err)
	}

	return err
}

// Get retrieves data from the database for the given user ID, phrase ID and mime type with pagination.
func (r *Repo) Get(userID, phraseID, limit, offset int64, mimeType entityuserphrase.MimeType) ([]entityuserphrase.UserPhrase, error) {
	var userPhrases []entityuserphrase.UserPhrase

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.
		Select("user_id", "phrase_id", "file_path", "mime_type").
		From("audio_retrieval_trx_user_phrases").
		Where(squirrel.Eq{"user_id": userID, "phrase_id": phraseID, "mime_type": mimeType}).
		OrderBy("id DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %v", err)
	}

	rows, err := r.dbInterface.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var userPhrase entityuserphrase.UserPhrase
		if err := rows.Scan(&userPhrase.UserID, &userPhrase.PhraseID, &userPhrase.FilePath, &userPhrase.MimeType); err != nil {
			return nil, fmt.Errorf("failed to scan row into struct: %v", err)
		}
		userPhrases = append(userPhrases, userPhrase)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %v", err)
	}

	return userPhrases, nil
}
