package User

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//json.unmarshal is used to convert byte array data into struct field......

type Profile struct {
	Name 	string	`json:"name"`
	UserId  int		`json:"user_id"`
	Desc    string  `json:"desc"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId,err := strconv.Atoi(vars["user_id"])
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("error here ...")
	}

	var  desc,name string
	iter := cql.Session.Query("select about_me,first_name from user.profile where id = ?", userId).Iter()
	iter.Scan(&desc, &name)

	up := Profile{
		Name:   	name,
		UserId:  	userId,
		Desc: 		desc,
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(up)
	if err != nil {
		log.Fatal("error encoding ")
	}

}
