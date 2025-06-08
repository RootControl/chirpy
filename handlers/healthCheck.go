package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"alive": true}`))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
