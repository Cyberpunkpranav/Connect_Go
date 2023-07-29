package controllers

import (
	sql "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Clinics_List(w http.ResponseWriter, _ *http.Request) {
	db := sql.DB
	var query string
	query = "SELECT * from clinic "
	ClinicDB, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var Clinics []structs.Clinics_struct
	for ClinicDB.Next() {
		var clinic structs.Clinics_struct
		ClinicDB.Scan(
			&clinic.Id,
			&clinic.Title,
			&clinic.Address,
			&clinic.Total_roooms,
			&clinic.Phone_number,
			&clinic.GSTIN,
			&clinic.State_code,
			&clinic.Ip_address_list,
			&clinic.Latitude,
			&clinic.Longitude,
			&clinic.Created_at,
			&clinic.Updated_at,
			&clinic.Created_by,
			&clinic.Updated_by,
			&clinic.Status,
		)
		query = fmt.Sprintf("SELECT * from clinic_rooms WHERE clinic_id ='%v'", *clinic.Id)
		RoomsDB, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		for RoomsDB.Next() {
			var room structs.Rooms_struct
			RoomsDB.Scan(
				&room.Id,
				&room.Clinic_id,
				&room.Room_number,
				&room.Room_type,
				&room.Map_image,
				&room.Created_at,
				&room.Updated_at,
				&room.Created_by,
				&room.Updated_by,
				&room.Status,
			)
			clinic.Clinic_Rooms = append(clinic.Clinic_Rooms, room)
		}
		Clinics = append(Clinics, clinic)
	}
	JSONAPI := structs.CLinic_List{
		Message: "Clinics List",
		Status:  true,
		Data:    Clinics,
	}
	jsondata, err := json.Marshal(JSONAPI)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsondata)
}
