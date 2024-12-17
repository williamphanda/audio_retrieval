package audio

import (
	entityuserphrase "audioretrieval/internal/entity/userphrase"
	"mime/multipart"
)

//go:generate mockgen -build_flags=-mod=mod -destination=handler.audio.dependencies_mock.go -package=audio -source=handler.audio.dependencies.go

type usecaseInterface interface {
	ProcessAudioFile(file multipart.File, fileHeader *multipart.FileHeader, userID, phraseID int64, mimeType entityuserphrase.MimeType) (string, error)
	CleanUpFile(filePath string) error
	RetrieveAudioFile(userID, phraseID int64, mimeType entityuserphrase.MimeType) (string, error)
}
