package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"quotation/quotation"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/download-quotation", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset-utf-8")
		query := r.URL.Query()
		path := query.Get("path")
		if path != "" {
			if err := quotation.RunQuotation(filepath.FromSlash(path)); err != nil {
				w.Write([]byte(`{"error":"Error to download quotation"}`))
				return
			}
			w.Write([]byte(fmt.Sprintf(`{"success":"Download file with successfully in %s"}`, path)))
			return
		}
	})

	http.ListenAndServe(":"+port, nil)
}
