package capture

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/FelipeSoft/go-secure-backup/internal/agent/entity"
	"github.com/FelipeSoft/go-secure-backup/internal/agent/platform"
)

type CaptureFilesAndFolders struct {
	storage entity.Storage
}

type BackupDTO struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	UserId string `json:"userId"`
}

func NewCaptureFilesAndFolders(storage entity.Storage) *CaptureFilesAndFolders {
	return &CaptureFilesAndFolders{
		storage: storage,
	}
}

func (s *CaptureFilesAndFolders) Execute() {
	operationalSystem := runtime.GOOS
	currentUserId := os.Getenv("USER_ID")
	res, err := http.Get(fmt.Sprintf("http://localhost:4816/backup/find/%s", currentUserId))
	if err != nil {
		log.Fatalf("HTTP Request failed: %s", err.Error())
	}

	var bkpRes BackupDTO
	err = json.NewDecoder(res.Body).Decode(&bkpRes)

	if err != nil {
		log.Fatalf("Decode Error: %s", err.Error())
	}

	content := platform.PlatformFactory(operationalSystem).GetContentFromPath(bkpRes.Path)
	for c := 0; c < len(content); c++ {
		s.storage.PutFile(content[c])
	}
}
