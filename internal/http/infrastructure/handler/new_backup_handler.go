package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
)

type NewBackup struct {
	NewBackupUseCase *usecase.NewBackup
}

func NewBackupHandler(newBackupUseCase *usecase.NewBackup) *NewBackup {
	return &NewBackup{
		NewBackupUseCase: newBackupUseCase,
	}
}

func (h *NewBackup) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Cannot %s %s", strings.ToUpper(r.Method), r.URL.Path)))
		return
	}

	var input usecase.NewBackupDTO
	body := r.Body
	err := json.NewDecoder(body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}
	defer body.Close()

	err = h.NewBackupUseCase.Execute(input)
	if err != nil {
		if err.Error() == "user backup already exists" {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
