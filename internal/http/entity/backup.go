package entity

type Backup struct {
	Id   string
	Path string
	User User
}

func (b *Backup) newBackup (id, path string, user User) *Backup {
	return &Backup{
		Id: id,
		Path: path,
		User: user,
	}
}
