package crud

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"

	"example.com/initdb"
)

type PetInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateSession() (session *gocql.Session) {
	session = initdb.Connect()
	return
}

func CloseSession(session *gocql.Session) {
	session.Close()
}

func Find(session *gocql.Session, id int) []map[string]interface{} {
	query := session.Query(`SELECT * FROM heartrate_v1 where id = ?`, id)
	rows, _ := query.Iter().SliceMap()
	return rows
}

func Insert(session *gocql.Session, pet map[string]interface{}) {
	err := session.Query(`INSERT INTO heartrate_v1 (id, name) VALUES (?, ?)`, pet["id"], pet["name"]).Exec()
	if err != nil {
		panic("Failed to insert")
	}
}

func Update(session *gocql.Session, pet map[string]interface{}) {
	err := session.Query(`UPDATE heartrate_v1 SET name = ? WHERE id = ?`, pet["name"], pet["id"]).Exec()
	if err != nil {
		panic(err)
	}
}

func Delete(session *gocql.Session, pet map[string]interface{}) {
	err := session.Query(`DELETE FROM heartrate_v1 WHERE id = ?`, pet["id"]).Exec()
	if err != nil {
		panic(err)
	}
}

func GetAndPrint(session *gocql.Session) (query *gocql.Query) {
	query = session.Query("SELECT * FROM heartrate_v1")
	printData(query)
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
			fmt.Printf("Id: %v, Name: %v\n", row["id"], row["name"])
		}
	} else {
		panic("Query error: " + err.Error())
	}
}
