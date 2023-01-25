package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Index = newIndexHandler() 

func newIndexHandler() http.Handler {
	router := mux.NewRouter()
	router.Use(commonMiddleware)


	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, access_token, device_id, x-diug-ofni")
		next.ServeHTTP(w, r)
	})

}