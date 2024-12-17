package audio

import (
	entityuserphrase "audioretrieval/internal/entity/userphrase"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

// ProcessAudioFile validate data, store, convert on both disk and database
func (u *Usecase) ProcessAudioFile(file multipart.File, fileHeader *multipart.FileHeader, userID, phraseID int64, mimeType entityuserphrase.MimeType) (string, error) {
	// validate user exist
	_, err := u.repoUser.GetByID(userID)
	if err != nil {
		return "", err
	}

	// validate phrase valid
	_, err = u.repoPhrase.GetByID(phraseID)
	if err != nil {
		return "", err
	}

	// create file path
	filePath := filepath.Join(
		"files/audio",
		strconv.FormatInt(userID, 10),
		strconv.FormatInt(phraseID, 10),
		strconv.FormatInt(time.Now().Unix(), 10),
		fileHeader.Filename,
	)

	// store original file, could be for audit purpose, or legal purpose, or other purposes
	// if we need to delete the original file, we could do so in async but it is not handled in current code
	err = u.repoFile.Store(file, filePath)
	if err != nil {
		return filePath, err
	}

	// TODO: if need more convertion, could add more convert method in goroutine
	// or if we choose to have a async ux, then we could trigger file conversion
	// in the background by using message queue. Currently it only support 1 type of convertion to OGG
	//
	// TODO: we could add audio lifecycle and state engine as well if needed
	// having them, for example, could help us to check that an audio file is currently converting
	convertedFilePath, err := u.repoFile.Convert(filePath)
	if err != nil {
		return convertedFilePath, err
	}

	// I assume user could upload multiple time of the user phrase
	// and we are complied to store the versioning for legal matters / training purposes
	err = u.repoUserPhrase.Insert(entityuserphrase.UserPhrase{
		UserID:   userID,
		PhraseID: phraseID,
		FilePath: filePath,
		MimeType: entityuserphrase.MimeTypeOGG, // TODO: since it only convert into 1 file type, then it is hardcoded here
	})
	return filePath, err
}

func (u *Usecase) CleanUpFile(filePath string) error {
	if filePath == "" {
		return nil
	}

	return u.repoFile.Delete(filePath)
}

// RetrieveAudioFile get user phrase file path and validate the file existance on disk
func (u *Usecase) RetrieveAudioFile(userID, phraseID int64, mimeType entityuserphrase.MimeType) (string, error) {
	const (
		limit  int64 = 1
		offset int64 = 0
	)

	audioFiles, err := u.repoUserPhrase.Get(userID, phraseID, limit, offset, mimeType)
	if err != nil {
		return "", err
	}

	if len(audioFiles) == 0 {
		return "", errors.New("record not found")
	}

	return u.repoFile.Retrieve(audioFiles[0].FilePath)
}
