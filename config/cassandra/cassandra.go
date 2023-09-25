package cassandra

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

type credential struct {
	CASSANDRA_HOST     string
	CASSANDRA_PORT     int
	CASSANDRA_USER     string
	CASSANDRA_PASSWORD string
	CASSANDRA_KEYSPACE string
}

func Connect() (*gocql.Session, error) {
	CASSANDRA_PORT, err := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
	if err != nil {
		log.Fatal("Error loading CASSANDRA_PORT from environment variable!")
	}

	cassandraCredential := credential{
		CASSANDRA_HOST:     os.Getenv("CASSANDRA_HOST"),
		CASSANDRA_PORT:     CASSANDRA_PORT,
		CASSANDRA_USER:     os.Getenv("CASSANDRA_USER"),
		CASSANDRA_PASSWORD: os.Getenv("CASSANDRA_PASSWORD"),
		CASSANDRA_KEYSPACE: os.Getenv("CASSANDRA_KEYSPACE"),
	}

	cluster := gocql.NewCluster(fmt.Sprintf("%s:%d", cassandraCredential.CASSANDRA_HOST, cassandraCredential.CASSANDRA_PORT))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraCredential.CASSANDRA_USER,
		Password: cassandraCredential.CASSANDRA_PASSWORD,
	}
	cluster.Keyspace = cassandraCredential.CASSANDRA_KEYSPACE

	Session, err = cluster.CreateSession()
	if err != nil {
		return Session, err
	}

	return Session, nil
}
