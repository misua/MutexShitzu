package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, " Category is : %v\n", vars["category"])
	fmt.Fprintf(w, "ID is : %v\n", vars["id"])

}

func Encodeverse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//id := vars["id"]
	fmt.Fprintf(w, "Category ID: %v\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.UseEncodedPath().SkipClean(true)
	//r.EncodedPathDelimiter(';')
	r.PathEncoder(func(r http.Request, path string) string {
		path = strings.ReplaceAll(path, "%2F", "/")
		return mux.EncodePathSegment(path)

	})
	//.NewRoute().Path("/category/id")
	// attach a path with handler

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		//good prace to enforce times for your servers

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.HandleFunc("/category/{id}", Encodeverse)

	//url, err := r.Get("articleRoute").URL("category", "books", "id", "123")
	//fmt.Printf(url.URL)

	log.Fatal(srv.ListenAndServe())

}
