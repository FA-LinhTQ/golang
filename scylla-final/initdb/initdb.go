package initdb

import (
	"github.com/gocql/gocql"
)

func Connect() *gocql.Session {
	var cluster = gocql.NewCluster("node-0.aws-ap-southeast-1.4ff680e7460310e78cb2.clusters.scylla.cloud")
	cluster.Keyspace = "pets_clinic"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "6hC9JEBD5wTNQdo"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")

	var session, err = cluster.CreateSession()
	if err != nil {
		panic("Failed to connect to cluster")
	}
	return session
}
