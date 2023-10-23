package models

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

var DB *gocql.Session

func ConnectDataBase() {

	host := os.Getenv("DB_HOSTS")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	cluster := gocql.NewCluster(host + ":" + port)
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = os.Getenv("DB_KEYSPACE")
	cluster.ProtoVersion = 4
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	sess, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	DB = sess
	log.Println("Successfully connected to Database")

}
