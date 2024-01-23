package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"

	"example.com/initdb"
)

type PetInfo struct {
	Name      string
	HeartRate int
	Time      time.Time
}

func main() {
	// Get a greeting message and print it.
	session := initdb.Connect()
	defer session.Close()
	pet := PetInfo{
		"Milu",
		20,
		time.Now(),
	}
	insert(session, pet)
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
	query = session.Query("SELECT * FROM heartrate_v1")
	printData(query)
	return
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
