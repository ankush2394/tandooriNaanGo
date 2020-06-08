package User

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"tandoorinaan/golang/tandoorinaan-api/Redis"
)

//json.unmarshal is used to convert byte array data into struct field......

type Profile struct {
	Name 	string	`json:"name"`
	UserId  int		`json:"user_id"`
	Desc    string  `json:"desc"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {

	var ctx context.Context
	redisClient := Redis.GetInstance()
	var userProfile Profile

	vars := mux.Vars(r)
	userId,err := strconv.Atoi(vars["user_id"])
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("invalid user id  ...")
	}

	userProfileKey := getUserProfileRedisKey(userId)
	res,err := redisClient.Get(ctx,userProfileKey).Result()
	if err != nil || strings.EqualFold(res,"") {
		userProfile = getUserProfileInfo(userId)
		redisClient.Set(ctx, userProfileKey, userProfile, 3600)

	} else {
		res, err := redisClient.Get(ctx, userProfileKey).Result()
		if err != nil {
			log.Info("error getting from redis..", err)
			w.WriteHeader(404)
			json.NewEncoder(w).Encode("we are trying to up the reids ...")
			return
		} else {
			err := json.Unmarshal([]byte(res), &userProfile)
			if err != nil {
				log.Info("error in Unmarshal..", err)
				w.WriteHeader(404)
				return
			}
		}
	}

	up := Profile{
		Name:   	userProfile.Name,
		UserId:  	userProfile.UserId,
		Desc: 		userProfile.Desc,
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(up)
	if err != nil {
		log.Fatal("error encoding ")
	}

}

func getUserProfileRedisKey(userId int) string {
	return fmt.Sprintf("user_profile_%s",strconv.Itoa(userId))
}

func getUserProfileInfo(userId int) Profile {

	var  desc,name string
	iter := cql.Session.Query("select about_me,first_name from user.profile where id = ?", userId).Iter()
	iter.Scan(&desc, &name)

	up := Profile{
		Name:   name,
		UserId: userId,
		Desc:   desc,
	}
	return up
}