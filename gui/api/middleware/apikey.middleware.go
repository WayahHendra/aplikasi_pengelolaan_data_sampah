package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"trash-app/core"
	"trash-app/utils"
)

func APIKeyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")

		validAPIKey := core.API_KEY
		if validAPIKey == "" {
			validAPIKey = "6f64299c-9e32-463d-a662-4800b73efe17" // fallback kalau belum di-set
		}

		if apiKey != validAPIKey {
			log.Print(utils.Types["error"], "Unauthorized: Invalid API Key")
			w.WriteHeader(http.StatusUnauthorized)

			response := map[string]interface{}{
				"status":  "error",
				"message": "Unauthorized: Invalid API Key",
			}

			json.NewEncoder(w).Encode(response)

			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}
