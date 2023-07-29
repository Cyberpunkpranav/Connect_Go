package controllers

import (
	sql "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"

	// "fmt"
	"log"
	"net/http"
)

func Vaccines_List(w http.ResponseWriter) {
	db := sql.DB
	vaccinestable, err := db.Query("SELECT id,name,vaccine_category_id FROM vaccines")
	if err != nil {
		log.Fatal(err)
	}
	var vaccine []structs.Vaccines_struct
	for vaccinestable.Next() {
		var vaccines structs.Vaccines_struct
		err := vaccinestable.Scan(&vaccines.Id, &vaccines.Vaccine_name, &vaccines.Vaccine_category_id)
		if err != nil {
			panic(err.Error())
		}
		vaccine = append(vaccine, vaccines)
	}
	if err := vaccinestable.Err(); err != nil {
		panic(err.Error())
	}
	jsonData, err := json.Marshal(vaccine)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	// fmt.Printf("%v", jsonData)
	w.Write(jsonData)
}
