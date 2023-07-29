package controllers

import (
	sql "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func All_Patients(w http.ResponseWriter, r *http.Request) {
	db := sql.DB
	var query string
	var (
		search = r.URL.Query().Get("search")
		limit  = r.URL.Query().Get("limit")
		offset = r.URL.Query().Get("offset")
	)
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	if search != "" {
		query = fmt.Sprintf("SELECT * FROM patients WHERE full_name LIKE '%s' LIMIT %s OFFSET %s", "%"+search+"%", limit, offset)
	}
	if search == "" {
		query = fmt.Sprintf("SELECT * FROM patients LIMIT %s OFFSET %s", limit, offset)
	}
	PatientsDB, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var All_Patients []structs.AllPatients
	for PatientsDB.Next() {
		var patient structs.AllPatients
		PatientsDB.Scan(
			&patient.Id,
			&patient.Full_name,
			&patient.Phone_country_code,
			&patient.Phone_number,
			&patient.Email,
			&patient.Gender,
			&patient.Pin_code,
			&patient.Fcm_token,
			&patient.Device_id,
			&patient.Version,
			&patient.Platform,
			&patient.Dob,
			&patient.Age,
			&patient.Relation,
			&patient.Link_id,
			&patient.Location,
			&patient.Latitude,
			&patient.Longitude,
			&patient.Membership_type_id,
			&patient.Source,
			&patient.Created_at,
			&patient.Updated_at,
			&patient.Created_by,
			&patient.Updated_by,
			&patient.Status,
		)
		query = fmt.Sprintf("SELECT * FROM patient_address WHERE patient_id = '%v'", *patient.Id)
		AddressDB, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		for AddressDB.Next() {
			var address structs.Address_struct
			AddressDB.Scan(
				&address.Id,
				&address.Patient_id,
				&address.Full_name,
				&address.Address_line1,
				&address.Address_line2,
				&address.Zip_code,
				&address.Country,
				&address.State,
				&address.City,
				&address.Created_at,
				&address.Updated_at,
				&address.Status,
			)
			patient.Address = append(patient.Address, address)
		}
		All_Patients = append(All_Patients, patient)
	}
	if search != "" {
		query = fmt.Sprintf("SELECT COUNT(id) AS total_count FROM patients WHERE full_name LIKE '%s' ", "%"+search+"%")
	}
	if search == "" {
		query = fmt.Sprintf("SELECT COUNT(id) AS total_count FROM patients")
	}

	var totalPatients int
	err = db.QueryRow(query).Scan(&totalPatients)
	if err != nil {
		log.Fatal(err)
	}
	patientsdata := structs.Patient_Data{
		Total_Count: totalPatients,
		Data:        All_Patients,
	}
	json_API := structs.Patient_List{
		Message: "Patients List",
		Status:  true,
		Data:    patientsdata,
	}
	jsonData, err := json.Marshal(json_API)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonData)
}
