package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
)

type FindBackupById struct {
	FindBackupByIdUseCase *usecase.FindBackupById
}

func NewFindBackupByIdHandler(findBackupByIdUseCase *usecase.FindBackupById) *FindBackupById {
	return &FindBackupById{
		FindBackupByIdUseCase: findBackupByIdUseCase,
	}
}

func (h *FindBackupById) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Cannot %s %s", r.Method, r.URL.Path)))
		return
	}

	backupId := r.PathValue("id")
	if backupId == "" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("backup id path value is required"))
		return
	}

	backup, err := h.FindBackupByIdUseCase.Execute(backupId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := json.Marshal(backup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
