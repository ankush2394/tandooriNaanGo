package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tandoorinaan/golang/tandoorinaan-api/User"
)


func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {


	r := mux.NewRouter()

	r.HandleFunc("/", heartbeat)
	r.HandleFunc("/user/{user_id}/profile", User.GetProfile)

	r.Use(loggingMiddleware)

	errHttp := http.ListenAndServe("localhost:8080", r)
	if errHttp != nil {
		log.Fatal("error connecting to localhost")
	}

}

type heartbeatResponse struct {
	Status		string		`json:"status"`
	Code 		int			`json:"code"`
}
func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}
