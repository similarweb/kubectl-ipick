package cmd

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	t.Run("valid", func(t *testing.T) {
		path := fmt.Sprintf("%s/testutils/file.txt", pwd)
		if !fileExists(path) {
			t.Fatalf("file %s should be found.", path)
		}

	})

	t.Run("invalid", func(t *testing.T) {
		path := fmt.Sprintf("%s/testutils/not-found.txt", pwd)
		if fileExists(path) {
			t.Fatalf("file %s should be not found.", path)
		}
	})
}

func TestFind(t *testing.T) {

	slice := []string{"a", "b", "c", "d"}

	t.Run("found", func(t *testing.T) {
		expectedStringFind := "a"
		_, found := find(slice, expectedStringFind)
		if !found {
			t.Fatalf("string `%s` should be found.", expectedStringFind)
		}
	})

	t.Run("not found", func(t *testing.T) {
		expectedStringFind := "e"
		_, found := find(slice, expectedStringFind)
		if found {
			t.Fatalf("string `%s` should be not found.", expectedStringFind)
		}
	})
}
