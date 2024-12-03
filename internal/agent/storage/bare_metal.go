package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/FelipeSoft/go-secure-backup/internal/agent/entity"
)

type BareMetal struct{}

func NewBareMetalStorage() *BareMetal {
	return &BareMetal{}
}

func (s *BareMetal) PutFile(content *entity.Content) {
	userId := os.Getenv("USER_ID")
	basePath := fmt.Sprintf("C:/bkp_bare_metal/%s", userId)

	_, err := os.Stat(basePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(basePath, 0750)
		if err != nil {
			log.Fatalf("Fail on base directory creation: %s", err.Error())
		}
	}

	path := filepath.Join(basePath, content.Path)
	err = os.MkdirAll(filepath.Dir(path), 0750)
	if err != nil {
		log.Fatalf("Fail on subdirectories creation: %s", err.Error())
	}

	err = os.WriteFile(path, content.Bytes, 0644)
	if err != nil {
		log.Fatalf("Fail on files creation: %s", err.Error())
	}
}
