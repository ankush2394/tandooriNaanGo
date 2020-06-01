package Cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"sync"
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
	cql.cluster = gocql.NewCluster("127.0.0.1") //contains list of cassandra machines in a cluster...
	cql.cluster.Consistency = gocql.One                //only one node in cluster
	cql.cluster.Timeout = 2000 * time.Millisecond
	cql.cluster.ConnectTimeout = 2000 * time.Millisecond
	cql.Session, err = cql.cluster.CreateSession()
	cql.cluster.ProtoVersion = 3

	//cql.cluster.Keyspace = cassandraKeySpace
	//defer cql.Session.Close()

	if err != nil {
		fmt.Println("error")
		log.Fatal("error in creating session of cassandra..", err)
	}

	return &cql
}
