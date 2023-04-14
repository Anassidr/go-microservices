package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local is an implementation of the Storage interface which works with the
// local disk on the current machine

type Local struct {
	maxFileSize int // maximum number of bytes for files
	basePath    string
}

func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &Local{basePath: p}, nil
}

// Save the contents of the Writer to the given path
// path is a relative path, basePath will be appended

func (l *Local) Save(path string, contents io.Reader) error {
	// get the full path for the file
	fp := l.fullPath(path)

	// get the directory and make sure it exists
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)
	}
}

// returns absolute path

func (l *Local) fullPath(path string) string {
	// append the given path to the base path
	return filepath.Join(l.basePath, path)
}