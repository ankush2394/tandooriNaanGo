package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"tandoorinaan/golang/tandoorinaan-api/Config/local"
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

	env := getEnvironment()
	appHost := env.Server.Host
	appPort := env.Server.Port
	errHttp := http.ListenAndServe(appHost+":"+strconv.Itoa(appPort), r)
	if errHttp != nil {
		log.Fatal("error connecting to "+ appHost)
	}

}

type heartbeatResponse struct {
	Status		string		`json:"status"`
	Code 		int			`json:"code"`
}
func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}

func getEnvironment() *local.Config{
	return local.Instance()
}
