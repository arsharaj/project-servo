package handlers

import (
	"net/http"

	"github.com/arsharaj/project-servo/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, "Service is healthy", map[string]string{
		"server_status": "OK",
	})
}
