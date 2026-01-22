package filestore

import "os"

type FileStorage interface {
	// Setup creates all necessary folders/files for storage
	Setup() error

	// Exists tells whether a file/directory path exists
	Exists(path string) (bool, error)

	// CreateAll attempts to create all folders until the provided path. Does nothing if folder already exists
	CreateAll(path string) error

	// ServeFile return the file on specified path, or null if not found
	ServeFile(path string) (*os.File, error)

	// UploadFile creates a file at given path and writes given list of bytes
	UploadFile(path string, bytes []byte) error

	// DeleteFile deletes a file at given path if it exists
	//
	// This is safe to use for files in directories that might not exist yet (in that case, will just return)
	DeleteFile(path string) error
}
