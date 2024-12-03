package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
)

type UpdateBackup struct {
	UpdateBackupUseCase *usecase.UpdateBackup
}

func NewUpdateBackupHandler(updateBackupUseCase *usecase.UpdateBackup) *UpdateBackup {
	return &UpdateBackup{
		UpdateBackupUseCase: updateBackupUseCase,
	}
}

func (h *UpdateBackup) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Cannot %s %s", strings.ToUpper(r.Method), r.URL.Path)))
		return
	}

	var input usecase.UpdateBackupDTO
	body := r.Body
	err := json.NewDecoder(body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Body"))
		return
	}

	input.Id = r.PathValue("id")
	err = h.UpdateBackupUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}