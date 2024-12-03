package usecase

import "github.com/FelipeSoft/go-secure-backup/internal/http/entity"

type DeleteBackup struct {
	repo entity.BackupRepository
}

func NewDeleteBackupUseCase(repo entity.BackupRepository) *DeleteBackup {
	return &DeleteBackup{
		repo: repo,
	}
}

func (uc *DeleteBackup) Execute(id string) error {
	if err := uc.repo.Delete(id); err != nil {
		return err
	}
	return nil
}