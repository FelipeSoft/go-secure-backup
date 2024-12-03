package entity

type Backup struct {
	Id     string
	Path   string
	UserId string
}

type BackupRepository interface {
	Create(backup Backup) error
	FindAll() ([]*Backup, error)
	FindById(id string) (*Backup, error)
	Update(backup Backup) error
	Delete(id string) error
}

func (b *Backup) newBackup(id, path string, userId string) *Backup {
	return &Backup{
		Id:   id,
		Path: path,
		UserId: userId,
	}
}
