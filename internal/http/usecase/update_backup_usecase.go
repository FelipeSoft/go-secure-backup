package usecase

import (
	"errors"
	"fmt"

	"github.com/FelipeSoft/go-secure-backup/internal/http/entity"
)

type UpdateBackup struct {
	repo entity.BackupRepository
}

type UpdateBackupDTO struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	UserId string `json:"userId"`
}

func NewUpdateBackupUseCase(repo entity.BackupRepository) *UpdateBackup {
	return &UpdateBackup{
		repo: repo,
	}
}

func (uc *UpdateBackup) Execute(input UpdateBackupDTO) error {
	if input.Id == "" {
		return errors.New("backup id is required")
	}
	if input.Path == "" && input.UserId == "" {
		return errors.New("path or userId backup are required")
	}

	foundBackup, err := uc.repo.FindById(input.Id)
	if err != nil {
		return err
	}

	if foundBackup == nil {
		return fmt.Errorf("cannot find backup with provided id %s", input.Id)
	}

	backup := entity.Backup{
		Id:     foundBackup.Id,
		Path:   foundBackup.Path,
		UserId: foundBackup.UserId,
	}

	if input.Path != "" {
		backup.Path = input.Path
	} else {
		backup.Path = foundBackup.Path
	}

	if input.UserId != "" {
		backup.UserId = input.UserId
	} else {
		backup.UserId = foundBackup.UserId
	}

	if err = uc.repo.Update(backup); err != nil {
		return err
	}
	
	return nil
}
