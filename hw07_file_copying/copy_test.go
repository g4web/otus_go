package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("Copy a non-existent file", func(t *testing.T) {
		err := Copy("non-existent-file.txt", "testdata/newFile.txt", 0, 0)
		require.NotNil(t, err)
	})

	t.Run("Copy a non-existent folder", func(t *testing.T) {
		err := Copy("./non-existent-folder/", "testdata/newFile.txt", 0, 0)
		require.NotNil(t, err)
	})

	t.Run("Copy an empty file", func(t *testing.T) {
		err := Copy("testdata/empty_file.txt", "testdata/newFile.txt", 0, 0)
		require.NotNil(t, err)
	})

	t.Run("Copy a folder", func(t *testing.T) {
		err := Copy("testdata/", "testdata/newFile.txt", 0, 0)
		require.NotNil(t, err)
	})

	t.Run("The offset is larger than the file size", func(t *testing.T) {
		err := Copy("testdata/out_offset0_limit10.txt", "testdata/newFile.txt", 11, 0)
		require.NotNil(t, err)
	})

	t.Run("Offset 0, limit 0", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset0_limit0.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 0, 0)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})

	t.Run("Offset 0, limit 10", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset0_limit10.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 0, 10)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		require.Nil(t, err)

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})

	t.Run("Offset 0, limit 1000", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset0_limit1000.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 0, 1000)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		require.Nil(t, err)

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})

	t.Run("Offset 0, limit 10000", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset0_limit10000.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 0, 10000)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		require.Nil(t, err)

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})

	t.Run("Offset 100, limit 1000", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset100_limit1000.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 100, 1000)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		require.Nil(t, err)

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})

	t.Run("Offset 6000, limit 1000", func(t *testing.T) {
		referenceFilePath := "testdata/out_offset6000_limit1000.txt"
		newFilePath := referenceFilePath + "_temp"

		err := Copy("testdata/input.txt", newFilePath, 6000, 1000)
		if err == nil {
			defer os.Remove(newFilePath)
		}

		require.Nil(t, err)

		newFile, _ := os.Open(newFilePath)
		newFileInfo, _ := newFile.Stat()

		referenceFile, _ := os.Open(referenceFilePath)
		referenceFileInfo, _ := referenceFile.Stat()

		require.Equal(t, newFileInfo.Size(), referenceFileInfo.Size())
	})
}
