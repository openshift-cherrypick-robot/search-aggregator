package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

// Global status of all clusters.
type status struct {
	TotalClusters int
	Message       string
}

// GetStatus responds with the global status.
func GetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	glog.Info("GetStatus() - TODO: Respond with all clusters and their last sync time and current hash.")
	var status = status{
		Message:       "TODO: This will respond with all clusters and their last sync time and current hash.",
		TotalClusters: 99, // TODO: Get total clusters from Redis
	}
	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		glog.Error("Error responding to GetStatus", err, status)
	}
}