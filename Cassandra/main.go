package Cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
	"sync"
	"tandoorinaan/golang/tandoorinaan-api/Config/local"
	"time"
)

var (
	once sync.Once
	cql  *Cassandra
)

type Cassandra struct {
	cluster    *gocql.ClusterConfig
	Session    *gocql.Session
}

func Instance() *Cassandra {
	once.Do(func() {
		cql = NewCqlConnection()
	})
	return cql
}


func NewCqlConnection() *Cassandra {

	var err error
	cql := Cassandra{}

	localEnv := getEnvironment()
	clusterHostFromConfig := localEnv.Database.Host
	cql.cluster = gocql.NewCluster(clusterHostFromConfig)  //contains list of cassandra machines in a cluster...
	cql.cluster.Consistency = gocql.One                   //only one node in cluster
	cql.cluster.Timeout = 2000 * time.Millisecond
	cql.cluster.ConnectTimeout = 2000 * time.Millisecond
	cql.Session, err = cql.cluster.CreateSession()
	cql.cluster.ProtoVersion = 3
	log.Info("cassandra address is ", clusterHostFromConfig)

	if err != nil {
		fmt.Println("error")
		log.Fatal("error in creating session of cassandra..", err)
	}

	return &cql
}

func getEnvironment() *local.Config{
	return local.Instance()
}
