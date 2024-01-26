package initdb

import (
	"github.com/gocql/gocql"
)

func Connect() *gocql.Session {
	var cluster = gocql.NewCluster("node-0.aws-ap-southeast-1.23aab8f5507a9044d1d0.clusters.scylla.cloud", "node-1.aws-ap-southeast-1.23aab8f5507a9044d1d0.clusters.scylla.cloud", "node-2.aws-ap-southeast-1.23aab8f5507a9044d1d0.clusters.scylla.cloud")
	cluster.Keyspace = "pets_clinic"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "nir6Js1GxlVaX4C"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_AP_SOUTHEAST_1")

	var session, err = cluster.CreateSession()
	if err != nil {
		panic("Failed to connect to cluster")
	}
	return session
}
