package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/FelipeSoft/go-secure-backup/internal/http/infrastructure/handler"
	"github.com/FelipeSoft/go-secure-backup/internal/http/infrastructure/repository/mysql"
	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("[Environment File Error]: %s", err.Error())
	}

	r := http.NewServeMux()

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/go-secure-backup")
	if err != nil {
		log.Fatalf("[MySQL Connection Error]: %s", err.Error())
	}

	backupRepository := mysql.NewBackupRepositoryMySQL(db)

	// Find All
	findAllBackupsUseCase := usecase.NewFindAllBackupsUseCase(backupRepository)
	findAllBackupsHandler := handler.NewFindAllBackups(findAllBackupsUseCase)
	r.HandleFunc("/backup/all", findAllBackupsHandler.Execute)

	// Find By Id
	findBackupByIdUseCase := usecase.NewFindBackupById(backupRepository)
	findBackupByIdHandler := handler.NewFindBackupByIdHandler(findBackupByIdUseCase)
	r.HandleFunc("/backup/find/{id}", findBackupByIdHandler.Execute)

	// New Backup
	newBackupUseCase := usecase.NewBackupUseCase(backupRepository)
	newBackupHandler := handler.NewBackupHandler(newBackupUseCase)
	r.HandleFunc("/backup/create", newBackupHandler.Execute)

	// Update Backup
	updateBackupUseCase := usecase.NewUpdateBackupUseCase(backupRepository)
	updateBackupHandler := handler.NewUpdateBackupHandler(updateBackupUseCase)
	r.HandleFunc("/backup/update/{id}", updateBackupHandler.Execute)

	// Delete Backup
	deleteBackupUseCase := usecase.NewDeleteBackupUseCase(backupRepository)
	deleteBackupHandler := handler.NewDeleteBackupHandler(deleteBackupUseCase)
	r.HandleFunc("/backup/delete/{id}", deleteBackupHandler.Execute)

	host := fmt.Sprintf("127.0.0.1:%s", os.Getenv("HTTP_SERVER_PORT")) 
	fmt.Println(host)

	err = http.ListenAndServe(host, r)
	if err != nil {
		log.Fatalf("[Server Listening Error]: %s", err.Error())
	}
}