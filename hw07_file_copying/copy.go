package main

import (
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileReader, limit, err := createFileReader(fromPath, limit, offset)
	if err != nil {
		return err
	}
	defer fileReader.Close()

	fileWriter, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileWriter.Close()

	return copyFromReaderToWriter(fileReader, fileWriter, limit)
}

func createFileReader(fromPath string, limit int64, offset int64) (*os.File, int64, error) {
	fileReader, err := os.Open(fromPath)
	if err != nil {
		return nil, 0, ErrUnsupportedFile
	}

	fromFileInfo, err := fileReader.Stat()
	if err != nil {
		return nil, 0, err
	}

	if (fromFileInfo.Size() == 0) || (fromFileInfo.IsDir()) {
		return nil, 0, ErrUnsupportedFile
	}

	if offset > fromFileInfo.Size() {
		return nil, 0, ErrOffsetExceedsFileSize
	}

	if offset > 0 {
		_, err = fileReader.Seek(offset, io.SeekStart)
		if err != nil {
			return nil, 0, err
		}
	}

	leftover := fromFileInfo.Size() - offset
	if (limit == 0) || (limit > leftover) {
		limit = leftover
	}

	return fileReader, limit, nil
}

func copyFromReaderToWriter(fileReader *os.File, fileWriter *os.File, limit int64) error {
	bar := progressbar.DefaultBytes(limit, "Copying...")

	if limit > 0 {
		_, err := io.CopyN(io.MultiWriter(fileWriter, bar), fileReader, limit)
		if err != nil {
			return err
		}
		return nil
	}

	_, err := io.Copy(io.MultiWriter(fileWriter, bar), fileReader)
	if err != nil {
		return err
	}
	return nil
}
