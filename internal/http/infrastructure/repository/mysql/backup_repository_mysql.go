package mysql

import (
	"database/sql"
	"github.com/FelipeSoft/go-secure-backup/internal/http/entity"
)

type BackupRepositoryMySQL struct {
	db *sql.DB
}

func NewBackupRepositoryMySQL(db *sql.DB) *BackupRepositoryMySQL {
	return &BackupRepositoryMySQL{
		db: db,
	}
}

func (r *BackupRepositoryMySQL) Create(backup entity.Backup) error {
	query := "INSERT INTO backup (path, user_id) VALUES (?, ?)"
	rows, err := r.db.Query(query, backup.Path, backup.UserId)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (r *BackupRepositoryMySQL) FindAll() ([]*entity.Backup, error) {
	query := "SELECT id, path, user_id FROM backup"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var backups []*entity.Backup
	for rows.Next() {
		var backup entity.Backup
		err := rows.Scan(&backup.Id, &backup.Path, &backup.UserId)
		if err != nil {
			return nil, err
		}
		backups = append(backups, &backup)
	}

	return backups, nil
}

func (r *BackupRepositoryMySQL) FindById(id string) (*entity.Backup, error) {
	query := "SELECT id, path, user_id FROM backup WHERE id = ?"
	rows := r.db.QueryRow(query, id)
	var backup entity.Backup
	err := rows.Scan(&backup.Id, &backup.Path, &backup.UserId)
	if err != nil {
		return nil, err
	}
	return &backup, nil
}

func (r *BackupRepositoryMySQL) Update(backup entity.Backup) error {
	query := "UPDATE backup SET path = ?, user_id = ? WHERE id = ?"
	rows, err := r.db.Query(query, backup.Path, backup.UserId, backup.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *BackupRepositoryMySQL) Delete(id string) error {
	query := "DELETE FROM backup WHERE id = ?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
