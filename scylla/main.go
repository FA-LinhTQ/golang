package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// pet := PetInfo{
	// 	"Panda",
	// 	11,
	// 	time.Now(),
	// }
	session := initDB()
	defer session.Close()
	// insert(session, pet)
	// update(session)
	// delete(session)
	getAndPrint(session)

}

func insert(session *gocql.Session, pet PetInfo) {
	err := session.Query(`INSERT INTO heartrate_v1 (id, time, heart_rate, name) VALUES (?, ?, ?, ?)`,
		gocql.TimeUUID(), pet.Time, pet.HeartRate, pet.Name).Exec()
	if err != nil {
		panic("Failed to insert")
	}
}

func update(session *gocql.Session) {
	err := session.Query(`UPDATE heartrate_v2 SET name = ? WHERE id = ?`, "ngon", "2a5c87dd-b934-11ee-96b9-d8c4972ed04b").Exec()
	if err != nil {
		panic(err)
	}
}

func delete(session *gocql.Session) {
	err := session.Query(`DELETE FROM heartrate_v2 WHERE id = ?`, "2a5c87dd-b934-11ee-96b9-d8c4972ed04b").Exec()
	if err != nil {
		panic(err)
	}
}

func getAndPrint(session *gocql.Session) (query *gocql.Query) {
	query = session.Query("SELECT * FROM heartrate_v2")
	printData(query)
	return
}

func initDB() *gocql.Session {
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

func printData(query *gocql.Query) {
	if rows, err := query.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			// fmt.Printf("Name: %v, Heart Rate: %v, Time: %v\n", row["name"], row["heart_rate"], row["time"])
			fmt.Printf("Id: %v, Name: %v, Heart Rate: %v, Time: %v\n", row["id"], row["name"], row["heart_rate"], row["time"])
		}
	} else {
		panic("Query error: " + err.Error())
	}
}

type PetInfo struct {
	Name      string
	HeartRate int
	Time      time.Time
}
