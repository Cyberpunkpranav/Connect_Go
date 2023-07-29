package controllers

import (
	db "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Doctor_Wise_Patient_appointments(w http.ResponseWriter, r *http.Request) {
	db := db.DB
	var query string
	var (
		search = r.URL.Query().Get("search")
	)
	var patient_wise_doctors_struct structs.PatientWiseDoctors_struct
	var patients []structs.Patient

	if search != "" {
		query = fmt.Sprintf("SELECT id as patient_id , full_name,phone_number FROM patients WHERE full_name LIKE '%s' LIMIT 5 OFFSET 0", "%"+search+"%")
		PatientsDB, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer PatientsDB.Close()
		for PatientsDB.Next() {
			var patient structs.Patient
			err := PatientsDB.Scan(
				&patient.Patient_id,
				&patient.Full_name,
				&patient.Phone_number,
			)
			if err != nil {
				log.Fatal(err)
			}
			if len(patient.Doctor) == 0 {
				query = "SELECT id as doctor_id , doctor_name FROM doctor"
				DoctorsDB, err := db.Query(query)
				if err != nil {
					log.Fatal(err)
				}
				defer DoctorsDB.Close()
				for DoctorsDB.Next() {
					var doctor structs.Doc
					err := DoctorsDB.Scan(
						&doctor.Doctor_id,
						&doctor.Doctor_name,
					)
					if err != nil {
						log.Fatal(err)
					}
					query = fmt.Sprintf("SELECT id as Appointment_id ,appointment_date from appointments where patient_id =%v AND doctor_id = %v ", *patient.Patient_id, *doctor.Doctor_id)
					appointmentDB, err := db.Query(query)
					if err != nil {
						log.Fatal(err)
					}
					defer appointmentDB.Close()
					for appointmentDB.Next() {
						var appointment structs.Appointment
						err := appointmentDB.Scan(
							&appointment.Appointment_id,
							&appointment.Appointment_date,
						)
						if err != nil {
							log.Fatal(err)
						}
						doctor.Appointments = append(doctor.Appointments, appointment)
					}
					patient.Doctor = append(patient.Doctor, doctor)

				}
			}
			patients = append(patients, patient)
		}
		patient_wise_doctors_struct.Patient = patients
		var apidata = structs.Appointments_Count{
			Message: "Patient-wise Doctor Appointments",
			Status:  true,
			Data:    patient_wise_doctors_struct,
		}
		jsonData, err := json.Marshal(apidata)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonData)
	} else {
		var apidata = structs.Appointments_Count{
			Message: "Search for a patient to get the data ",
			Status:  false,
			Data:    patient_wise_doctors_struct,
		}
		jsonData, err := json.Marshal(apidata)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonData)
	}
	// stats := db.Stats()
	// fmt.Printf("%+v\n", stats)
}
 