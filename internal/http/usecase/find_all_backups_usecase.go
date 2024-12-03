package usecase

import "github.com/FelipeSoft/go-secure-backup/internal/http/entity"

type FindAllBackups struct {
	repo entity.BackupRepository
}

type FindAllBackupsDTO struct {
	Id     string `json:"id"`
	Path   string `json:"path"`
	UserId string `json:"userId"`
}

func NewFindAllBackupsUseCase(repo entity.BackupRepository) *FindAllBackups {
	return &FindAllBackups{
		repo: repo,
	}
}

func (uc *FindAllBackups) Execute() ([]*FindAllBackupsDTO, error) {
	var output []*FindAllBackupsDTO

	res, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}

	for b := 0; b < len(res); b++ {
		backup := &FindAllBackupsDTO{
			Id:     res[b].Id,
			Path:   res[b].Path,
			UserId: res[b].UserId,
		}
		output = append(output, backup)
	}

	return output, nil
}
