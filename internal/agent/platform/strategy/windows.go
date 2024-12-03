package platform

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/FelipeSoft/go-secure-backup/internal/agent/entity"
)

type WindowsStrategy struct{}

func (s *WindowsStrategy) GetContentFromPath(path string) []*entity.Content {
	var output []*entity.Content

	err := filepath.WalkDir(path, func(fullPath string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		bytes, err := os.ReadFile(fullPath)
		if err != nil {
			return err
		}

		var correctFullPath string = fullPath
		if strings.Contains(fullPath, ":") {
			diskAndPath := strings.Split(fullPath, ":")
			disk := diskAndPath[0]
			path := diskAndPath[1]
			correctFullPath = fmt.Sprintf("%s_Drive%s", disk, path)
		}
		output = append(output, &entity.Content{
			Bytes: bytes,
			Path: correctFullPath,
		})
		return nil
	})

	if err != nil {
		log.Fatalf("error reading files: %s", err.Error())
	}

	return output
}
