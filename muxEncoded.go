package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.UseEncodedPath()

	// Custom path encoding function
	r.PathEncoder(func(r *http.Request, path string) string {
		// Replace %2F with /
		path = strings.ReplaceAll(path, "%2F", "/")
		return mux.EncodePathSegment(path)
	})

	r.HandleFunc("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprintf(w, "Category ID: %s", id)
	})

	http.ListenAndServe(":8080", r)
}
