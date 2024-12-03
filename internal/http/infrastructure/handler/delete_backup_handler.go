package handler

import (
	"fmt"
	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
	"net/http"
)

type DeleteBackup struct {
	deleteBackupUseCase *usecase.DeleteBackup
}

func NewDeleteBackupHandler(deleteBackupUseCase *usecase.DeleteBackup) *DeleteBackup {
	return &DeleteBackup{
		deleteBackupUseCase: deleteBackupUseCase,
	}
}

func (h *DeleteBackup) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Cannot %s %s", r.Method, r.URL.Path)))
		return
	}

	backupId := r.PathValue("id")

	if backupId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id path value is required"))
		return
	}

	if err := h.deleteBackupUseCase.Execute(backupId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
