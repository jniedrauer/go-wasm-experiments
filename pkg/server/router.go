package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

const assetPrefix = "./assets"

func getRouter() *mux.Router {
	r := mux.NewRouter()
	s := http.StripPrefix("/static/", http.FileServer(http.Dir(assetPrefix+"/static/")))
	r.PathPrefix("/static/").Handler(s)
	r.HandleFunc("/healthcheck", healthCheck)

	http.Handle("/", r)

	return r
}
