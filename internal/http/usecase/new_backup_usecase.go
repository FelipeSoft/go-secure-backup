package usecase

import (
	"errors"
	"github.com/FelipeSoft/go-secure-backup/internal/http/entity"
)

type NewBackup struct {
	repo entity.BackupRepository
}

type NewBackupDTO struct {
	Path   string `json:"path"`
	UserId string `json:"userId"`
}

func NewBackupUseCase(repo entity.BackupRepository) *NewBackup {
	return &NewBackup{
		repo: repo,
	}
}

func (uc *NewBackup) Execute(input NewBackupDTO) error {
	if input.Path == "" || input.UserId == "" {
		return errors.New("path and userId are required")
	}

	res, err := uc.repo.FindById(input.UserId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		backup := entity.Backup{
			Path: input.Path,
			UserId: input.UserId,
		}
		uc.repo.Create(backup)
	}

	if res != nil {
		return errors.New("user backup already exists")
	}

	return nil
}
