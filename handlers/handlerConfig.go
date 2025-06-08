package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type ApiConfig struct {
	ApiHits atomic.Int32
}

func NewApiConfig() *ApiConfig {
	return &ApiConfig{}
}

func (c *ApiConfig) IncrementApiHits(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.ApiHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (c *ApiConfig) GetApiHits() int32 {
	return c.ApiHits.Load()
}

func (c *ApiConfig) ResetApiHits() {
	c.ApiHits.Store(0)
}

func (c *ApiConfig) GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json := fmt.Sprintf(`{"hits": %d}`, c.GetApiHits())
	_, err := w.Write([]byte(json))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *ApiConfig) ResetMetrics(w http.ResponseWriter, r *http.Request) {
	c.ResetApiHits()
	w.WriteHeader(http.StatusOK)
}
