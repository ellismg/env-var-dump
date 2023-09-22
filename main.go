package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", os.Getenv("EVN_VAR_DUMP_SHARED_KEY")) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		env := map[string]string{}
		for _, e := range os.Environ() {
			parts := strings.SplitN(e, "=", 2)

			if parts[0] == "EVN_VAR_DUMP_SHARED_KEY" {
				continue
			}

			env[parts[0]] = parts[1]
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(env)
	})

	http.ListenAndServe(":8080", nil)
}
