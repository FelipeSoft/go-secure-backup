package platform

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type WindowsStrategy struct{}

func (s *WindowsStrategy) GetContentFromPath(path string) []string {
	var items []string

	fmt.Println("Hello from Windows Strategy")
	err := filepath.WalkDir(path, func(fullPath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		items = append(items, fullPath)
		return nil
	})

	if err != nil {
		log.Fatalf("error reading files: %s", err.Error())
	}

	return items
}
