package audio

import (
	entityphrase "audioretrieval/internal/entity/phrase"
	entityuser "audioretrieval/internal/entity/user"
	entityuserphrase "audioretrieval/internal/entity/userphrase"
	"mime/multipart"
)

type repoFile interface {
	Convert(filePath string) (string, error)
	Store(file multipart.File, filePath string) error
	Delete(filePath string) error
	Retrieve(filePath string) (string, error)
}

// TODO: add cache repo if needed, for example could utilize redis
type repoPhrase interface {
	GetByID(phraseID int64) (entityphrase.Phrase, error)
}

// TODO: add cache repo if needed, for example could utilize redis
type repoUser interface {
	GetByID(userID int64) (entityuser.User, error)
}

type repoUserPhrase interface {
	Insert(userPhrase entityuserphrase.UserPhrase) error
	Get(userID, phraseID, limit, offset int64, mimeType entityuserphrase.MimeType) ([]entityuserphrase.UserPhrase, error)
}
