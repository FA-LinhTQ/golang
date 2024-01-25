package crud

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"

	"example.com/initdb"
)

type PetInfo struct {
	Name      string
	HeartRate int
	Time      time.Time
}

func CreateSession() (session *gocql.Session) {
	session = initdb.Connect()
	return
}

func CloseSession(session *gocql.Session) {
	session.Close()
}

func Insert(session *gocql.Session, pet PetInfo) {
	err := session.Query(`INSERT INTO heartrate_v1 (id, time, heart_rate, name) VALUES (?, ?, ?, ?)`,
		gocql.TimeUUID(), pet.Time, pet.HeartRate, pet.Name).Exec()
	if err != nil {
		panic("Failed to insert")
	}
}

func Update(session *gocql.Session) {
	err := session.Query(`UPDATE heartrate_v2 SET name = ? WHERE id = ?`, "ngon", "2a5c87dd-b934-11ee-96b9-d8c4972ed04b").Exec()
	if err != nil {
		panic(err)
	}
}

func Delete(session *gocql.Session) {
	err := session.Query(`DELETE FROM heartrate_v2 WHERE id = ?`, "2a5c87dd-b934-11ee-96b9-d8c4972ed04b").Exec()
	if err != nil {
		panic(err)
	}
}

func GetAndPrint(session *gocql.Session) (query *gocql.Query) {
	query = session.Query("SELECT * FROM heartrate_v1")
	Test(query)
	return
}

func Test(query *gocql.Query) {
	pets := make(map[int]PetInfo)
	if rows, err := query.Iter().SliceMap(); err == nil {
		for i, row := range rows {
			pets[i] = PetInfo{}
			log.Println(row["name"])

		}
	} else {
		panic("Query error: " + err.Error())
	}
	fmt.Println(pets)
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
