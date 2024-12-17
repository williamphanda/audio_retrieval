package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

// Convert convert any audio file type to WAV format
// it require ffmpeg to be installed in the server
// TODO: need to setup ansible or any equivalent to create image with ffmpeg
// TODO: need to setup a more flexible output converter if needed
func (r *Repo) Convert(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("target file does not exist: %s", filePath)
	}

	// TODO: find library to trim extension
	extension := filepath.Ext(filePath)
	baseName := filePath[:len(filePath)-len(extension)]
	outputFilePath := baseName + ".ogg"

	cmd := exec.Command("ffmpeg", "-i", filePath, outputFilePath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error converting file to WAV: %v", err)
	}
	return outputFilePath, nil
}

// Store function store file from input to disk
func (r *Repo) Store(file multipart.File, filePath string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	tempFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		return fmt.Errorf("failed to save converted file: %v", err)
	}

	return nil
}

// Delete removes a file from the disk
func (r *Repo) Delete(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

// Retrieve validate file existance
func (r *Repo) Retrieve(filePath string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}
	return filePath, nil
}
