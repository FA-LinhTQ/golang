package main

import (
	"github.com/gocql/gocql"
)

func main() {
	var cluster = gocql.NewCluster("node-0.aws-us-east-1.463d0373fb0b66377e84.clusters.scylla.cloud", "node-1.aws-us-east-1.463d0373fb0b66377e84.clusters.scylla.cloud", "node-2.aws-us-east-1.463d0373fb0b66377e84.clusters.scylla.cloud")
	cluster.Keyspace = "pets_clinic"
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "scylla", Password: "nDplmYd0x94TKvL"}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("AWS_US_EAST_1")

	var session, err = cluster.CreateSession()
	if err != nil {
		panic("Failed to connect to cluster")
	}

	defer session.Close()

	// var query = session.Query("SELECT * FROM pets_clinic.heartrate_v4") // keyspace setted
	// var query = session.Query("SELECT * FROM heartrate_v4")

	// insert
	// errs := session.Query(`INSERT INTO heartrate_v4 (pet_chip_id, time, heart_rate, pet_name) VALUES (?, ?, ?, ?)`,
	// 	gocql.TimeUUID(), time.Now(), 99, "hehe").Exec()
	// fmt.Println(errs)

	// if rows, err := query.Iter().SliceMap(); err == nil {
	// 	for _, row := range rows {
	// 		fmt.Printf("Name: %v, Heart Rate: %v\n", row["pet_name"], row["heart_rate"])
	// 	}
	// } else {
	// 	panic("Query error: " + err.Error())
	// }
}
