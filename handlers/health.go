// Package handlers contains HTTP handlers for the Servo server.
package handlers

import (
	"net/http"

	"github.com/arsharaj/project-servo/utils"
)

// HealthCheck responds with a basic health status and a 200 OK status code.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, "Service is healthy", map[string]string{
		"server_status": "OK",
	})
}
