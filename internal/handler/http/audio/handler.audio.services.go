package audio

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	entityuserphrase "audioretrieval/internal/entity/userphrase"
)

// TODO: move MaxFileSize to a configurable file
// MaxFileSize define the maximum size of an audio file.
const MaxFileSize = 5 * 1024 * 1024

// StoreAudio will validate, accept, convert, and store audio file from a request
func (h *Handler) StoreAudio(w http.ResponseWriter, r *http.Request) {
	// TODO: add monitoring + tracer code

	userIDStr := chi.URLParam(r, "user_id")
	phraseIDStr := chi.URLParam(r, "phrase_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	phraseID, err := strconv.ParseInt(phraseIDStr, 10, 64)
	if err != nil || phraseID <= 0 {
		http.Error(w, "Invalid phrase ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(MaxFileSize); err != nil {
		http.Error(w, "Audio file size exceed limit", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("audio_file")
	if err != nil {
		http.Error(w, "Audio File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if fileHeader == nil {
		http.Error(w, "Header Audio File not found", http.StatusBadRequest)
		return
	}

	if fileHeader.Size > MaxFileSize {
		http.Error(w, "Audio file size exceed limit", http.StatusBadRequest)
		return
	}

	fileType, err := validateFileType(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	filePath, err := h.usecase.ProcessAudioFile(file, fileHeader, userID, phraseID, fileType)
	if err != nil {
		h.cleanUpFailInsertion(filePath)

		log.Println("Fail to process audio file: ", err)
		http.Error(w, "Process audio failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validateFileType(file multipart.File) (entityuserphrase.MimeType, error) {
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return 0, fmt.Errorf("Unable to read file: %v", err)
	}

	fileType := http.DetectContentType(buffer)
	if _, ok := entityuserphrase.MapMimeTypeToEnum[fileType]; !ok {
		return 0, fmt.Errorf("Invalid file type: %s", fileType)
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return 0, fmt.Errorf("Failed to reset file pointer: %v", err)
	}

	return entityuserphrase.MapMimeTypeToEnum[fileType], nil
}

// cleanUpFailInsertion try to delete file in the system, and log if fail
// TODO: move this function to cron sweep or publish to message queue for async process
func (h *Handler) cleanUpFailInsertion(filePath string) {
	cleanUpErr := h.usecase.CleanUpFile(filePath)
	if cleanUpErr != nil {
		log.Printf("Failed to clean up file (%s): %v", filePath, cleanUpErr)
	}
}

func (h *Handler) GetAudio(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "user_id")
	phraseIDStr := chi.URLParam(r, "phrase_id")
	audioFormat := chi.URLParam(r, "audio_format")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	phraseID, err := strconv.ParseInt(phraseIDStr, 10, 64)
	if err != nil || phraseID <= 0 {
		http.Error(w, "Invalid phrase ID", http.StatusBadRequest)
		return
	}

	if _, ok := entityuserphrase.MapMimeTypeToEnum[audioFormat]; !ok {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	filePath, err := h.usecase.RetrieveAudioFile(userID, phraseID, entityuserphrase.MapMimeTypeToEnum[audioFormat])
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
