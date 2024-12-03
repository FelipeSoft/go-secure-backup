package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FelipeSoft/go-secure-backup/internal/http/usecase"
)

type FindAllBackups struct {
	FindAllBackupsUseCase *usecase.FindAllBackups
}

func NewFindAllBackups(findAllBackupsUseCase *usecase.FindAllBackups) *FindAllBackups {
	return &FindAllBackups{
		FindAllBackupsUseCase: findAllBackupsUseCase,
	}
}

func (h *FindAllBackups) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Cannot %s %s", r.Method, r.URL.Path)))
		return
	}

	backups, err := h.FindAllBackupsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(backups)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
}
