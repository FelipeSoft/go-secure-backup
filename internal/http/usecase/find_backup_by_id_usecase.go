package usecase

import "github.com/FelipeSoft/go-secure-backup/internal/http/entity"

type FindBackupById struct {
	repo entity.BackupRepository
}

type FindBackupByIdDTO struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	UserId string `json:"userId"`
}

func NewFindBackupById(repo entity.BackupRepository) *FindBackupById {
	return &FindBackupById{
		repo: repo,
	}
}

func (uc *FindBackupById) Execute(id string) (*FindBackupByIdDTO, error) {
	backup, err := uc.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	output := FindBackupByIdDTO{
		Id: backup.Id,
		Path: backup.Path,
		UserId: backup.UserId,
	}
	return &output, nil
}
