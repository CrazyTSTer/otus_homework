package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer sourceFile.Close()

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == 0 || sourceFileInfo.IsDir() {
		return ErrUnsupportedFile
	} else if offset > sourceFileInfo.Size() {
		return ErrOffsetExceedsFileSize
	} else if limit == 0 || limit > sourceFileInfo.Size()-offset {
		limit = sourceFileInfo.Size() - offset
	}

	if _, err = sourceFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	destinationFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	if _, err = io.CopyN(destinationFile, sourceFile, limit); err != nil {
		return err
	}

	return nil
}
