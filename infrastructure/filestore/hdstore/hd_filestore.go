package hdstore

import (
	"fmt"
	"log"
	"os"
	"strings"
	"taskflow/domain/entities"
	"taskflow/infrastructure/filestore"
)

type hdFileStorage struct {
	storageFolder string
}

func NewHDFileStorage(config entities.Config) filestore.FileStorage {
	fileStorage := hdFileStorage{
		storageFolder: config.FileStorage.StorageFolder,
	}

	// Attempt to set up storage
	err := fileStorage.Setup()
	if err != nil {
		panic(fmt.Sprintf("failed to setup file storage: %v", err))
	}

	return fileStorage
}

func (h hdFileStorage) Setup() error {
	// Create storage folder if not found
	err := os.MkdirAll(h.storageFolder, os.ModePerm)
	if err != nil {
		log.Printf("error in [MkdirAll | StorageFolder]: %v", err)
		return fmt.Errorf("error in [MkdirAll | StorageFolder]: %v", err)
	}

	return nil
}

func (h hdFileStorage) Exists(path string) (bool, error) {
	// Add a leading slash if needed
	path = formatLeadingSlash(path)

	fullPath := fmt.Sprintf("%s%s", h.storageFolder, path)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		log.Printf("error in [Exists]: %v", err)
		return false, fmt.Errorf("error in [Exists]: %v", err)
	}
	return true, nil
}

func (h hdFileStorage) CreateAll(path string) error {
	// Add a leading slash if needed
	path = formatLeadingSlash(path)

	fullPath := fmt.Sprintf("%s%s", h.storageFolder, path)
	return os.MkdirAll(fullPath, os.ModePerm)
}

func (h hdFileStorage) ServeFile(path string) (*os.File, error) {
	// Add a leading slash if not there yet
	path = formatLeadingSlash(path)

	fullPath := fmt.Sprintf("%s%s", h.storageFolder, path)

	_, err := os.Stat(fullPath)
	if err != nil {
		return nil, os.ErrNotExist
	}

	file, err := os.Open(fullPath)
	if err != nil {
		log.Printf("error in [Open | %s]: %v", fullPath, err)
		return nil, fmt.Errorf("error in [Open | %s]: %v", fullPath, err)
	}

	return file, nil
}

func (h hdFileStorage) UploadFile(
	path string,
	bytes []byte,
) error {
	// Add a leading slash if not there yet
	path = formatLeadingSlash(path)

	// Create file
	fullPath := fmt.Sprintf("%s%s", h.storageFolder, path)
	file, err := os.Create(fullPath)
	if err != nil {
		log.Printf("error in [Create | %s]: %v", fullPath, err)
		return fmt.Errorf("error in [Create | %s]: %v", fullPath, err)
	}

	// Write bytes to file
	_, err = file.Write(bytes)
	if err != nil {
		log.Printf("error in [Write | %s]: %v", fullPath, err)
		return fmt.Errorf("error in [Write | %s]: %v", fullPath, err)
	}

	// Close file
	err = file.Close()
	if err != nil {
		log.Printf("error in [Close | %s]: %v", fullPath, err)
		return fmt.Errorf("error in [Close | %s]: %v", fullPath, err)
	}

	return nil
}

func (h hdFileStorage) DeleteFile(path string) error {
	// Add a leading slash if not there yet
	path = formatLeadingSlash(path)

	// Create file
	fullPath := fmt.Sprintf("%s%s", h.storageFolder, path)

	// Check if file exists
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Printf("error in [Stat | %s]: %v", fullPath, err)
		return fmt.Errorf("error in [Stat | %s]: %v", fullPath, err)
	}

	// Delete file
	err = os.Remove(fullPath)
	if err != nil {
		log.Printf("error in [Remove | %s]: %v", fullPath, err)
		return fmt.Errorf("error in [Remove | %s]: %v", fullPath, err)
	}

	return nil
}

func formatLeadingSlash(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}
