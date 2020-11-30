package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("Doesn't exist", func(t *testing.T) {
		err := Copy("test", "/tmp", 0, 0)
		require.NotNil(t, err)
	})

	t.Run("fileFrom is dir", func(t *testing.T) {
		err := Copy("/tmp", "test", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("fileTo is dir", func(t *testing.T) {
		_, err := os.Create("test")
		if err != nil {
			t.Error(err)
		}
		defer os.Remove("test")
		err = Copy("/tmp", "test", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("copy empty file", func(t *testing.T) {
		_, err := os.Create("test")
		if err != nil {
			t.Error(err)
		}
		defer os.Remove("test")
		err = Copy("test", "test1", 10, 0)

		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("offset > file size", func(t *testing.T) {
		file, err := os.Create("test")
		if err != nil {
			t.Error(err)
		}
		file.WriteString("test")
		defer os.Remove("test")
		err = Copy("test", "test1", 10, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

}
