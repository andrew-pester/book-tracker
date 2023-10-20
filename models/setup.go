package models

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

func ConnectDataBase() *gocql.Session {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	cluster := gocql.NewCluster(host + ":" + port)
	cluster.Keyspace = dbName
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Host: %s, Port: %s, Username: %s, Password: %s, Database Name: %s", host, port, username, password, dbName)
	return session

}
