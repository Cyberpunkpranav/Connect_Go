package controllers

import (
	sql "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Doctors_List(w http.ResponseWriter, r *http.Request) {
	var query string
	//----------------------------------Database---------------------------------------------------
	db := sql.DB
	//----------------------------------Database---------------------------------------------------
	var (
		clinicid = r.URL.Query().Get("clinic_id")
		search   = r.URL.Query().Get("search")
		limit    = r.URL.Query().Get("limit")
		offset   = r.URL.Query().Get("offset")
	)
	if limit == "" {
		limit = "10"
	}
	if offset == "" {
		offset = "0"
	}
	if search == "" && clinicid != "" {
		query = fmt.Sprintf("SELECT * FROM doctor join speciality on doctor.speciality_id = speciality.id WHERE clinic_id ='%s' LIMIT %s OFFSET %s", clinicid, limit, offset)
	} else if search == "" && clinicid == "" {
		query = fmt.Sprintf("SELECT * FROM doctor join speciality on doctor.speciality_id = speciality.id LIMIT %s OFFSET %s", limit, offset)
	} else if clinicid == "" && search != "" {
		query = fmt.Sprintf("SELECT * FROM doctor join speciality on doctor.speciality_id = speciality.id WHERE doctor_name LIKE '%s'", "%"+search+"%")
	} else {
		query = fmt.Sprintf("SELECT * FROM doctor join speciality on doctor.speciality_id = speciality.id WHERE doctor_name LIKE '%s' AND clinic_id ='%s' LIMIT %s OFFSET %s", "%"+search+"%", clinicid, limit, offset)

	}

	//---------------------------------Doctors List------------------------------------------------

	dataDB, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	Doctors := []structs.Doctor_info{}

	for dataDB.Next() {
		var doctor structs.Doctor_info
		err := dataDB.Scan(
			&doctor.Id,
			&doctor.User_id,
			&doctor.Clinic_id,
			&doctor.Room_id,
			&doctor.Doctor_name,
			&doctor.Email,
			&doctor.Phone_country_code,
			&doctor.Phone_number,
			&doctor.Password,
			&doctor.Fcm_token,
			&doctor.Gender,
			&doctor.Address,
			&doctor.City_id,
			&doctor.Languages,
			&doctor.Speciality_id,
			&doctor.Registration_num,
			&doctor.Medical_authority_id,
			&doctor.Degree_suffix,
			&doctor.Expertise_subSpecialty,
			&doctor.Image,
			&doctor.Pdf_header,
			&doctor.Pdf_footer,
			&doctor.ScheduleDays,
			&doctor.ScheduleTimings,
			&doctor.Hours_per_month,
			&doctor.Agreement_file,
			&doctor.Start_date,
			&doctor.Agreement_renew_date,
			&doctor.ConsulationFee,
			&doctor.Time_per_patient,
			&doctor.Appointment_phone_number,
			&doctor.Website,
			&doctor.Instagram,
			&doctor.Facebook,
			&doctor.Youtube,
			&doctor.Profile_share_count,
			&doctor.Profile_verified,
			&doctor.Profile_status,
			&doctor.Reset_token,
			&doctor.On_request,
			&doctor.Pre_pay_enrolled,
			&doctor.Pre_pay_discount,
			&doctor.Total_experience,
			&doctor.About,
			&doctor.Created_at,
			&doctor.Updated_at,
			&doctor.Created_by,
			&doctor.Updated_by,
			&doctor.Status,
			&doctor.Rent_amount,
			&doctor.Security_amount,
			&doctor.Speciality.Id,
			&doctor.Speciality.Name,
			&doctor.Speciality.Icon_url,
			&doctor.Speciality.Status,
			&doctor.Speciality.Description,
			&doctor.Speciality.Conditions,
			&doctor.Speciality.Created_at,
			&doctor.Speciality.Updated_at,
			&doctor.Speciality.Bg_color1,
			&doctor.Speciality.Bg_color2,
			&doctor.Speciality.Icon_color,
			&doctor.Speciality.Created_by,
			&doctor.Speciality.Updated_by,
		)
		if err != nil {
			log.Fatal(err)
		}
		query = fmt.Sprintf("Select * FROM doctor_timings where doctor_id=%v ", *doctor.Id)
		timingsDB, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		for timingsDB.Next() {
			var timing structs.Timings_struct
			err := timingsDB.Scan(
				&timing.Id,
				&timing.Doctor_id,
				&timing.ScheduleDays,
				&timing.ScheduleTimings,
				&timing.Time_per_patient,
				&timing.Clinic_id,
				&timing.Clinic_rooms_id,
				&timing.Created_by,
				&timing.Updated_by,
				&timing.Created_at,
				&timing.Updated_at,
				&timing.Status,
			)
			if err != nil {
				log.Fatal(err)
			}
			query = fmt.Sprintf("SELECT * FROM clinic where id = %v", *timing.Clinic_id)
			clinicDB, err := db.Query(query)
			if err != nil {
				log.Fatal(err)
			}
			for clinicDB.Next() {
				err := clinicDB.Scan(
					&timing.Clinic.Id,
					&timing.Clinic.Title,
					&timing.Clinic.Address,
					&timing.Clinic.Total_rooms,
					&timing.Clinic.Phone_number,
					&timing.Clinic.GSTIN,
					&timing.Clinic.State_code,
					&timing.Clinic.Ip_address_list,
					&timing.Clinic.Latitude,
					&timing.Clinic.Longitude,
					&timing.Clinic.Created_at,
					&timing.Clinic.Updated_at,
					&timing.Clinic.Created_by,
					&timing.Clinic.Updated_by,
					&timing.Clinic.Status,
				)
				if err != nil {
					log.Fatal(err)
				}
			}
			doctor.Timings = append(doctor.Timings, timing)
		}
		Doctors = append(Doctors, doctor)
	}

	//---------------------------------Doctors List------------------------------------------------

	//---------------------------------Doctors Count------------------------------------------------

	if search != "" && clinicid == "" {
		query = fmt.Sprintf("SELECT COUNT(id) AS total_count FROM doctor WHERE doctor_name LIKE '%s' ", "%"+search+"%")
	} else if search == "" && clinicid != "" {
		query = fmt.Sprintf("SELECT COUNT(id) AS total_count FROM doctor WHERE clinic_id ='%s' ", clinicid)
	} else if search != "" && clinicid != "" {
		query = fmt.Sprintf("SELECT COUNT(id) AS total_count FROM doctor WHERE doctor_name LIKE '%s' AND clinic_id ='%s' ", "%"+search+"%", clinicid)
	} else {
		query = "SELECT COUNT(id) AS total_count FROM doctor"
	}

	var totalDoctors int
	err = db.QueryRow(query).Scan(&totalDoctors)
	if err != nil {
		log.Fatal(err)
	}

	//---------------------------------Doctors Count-------------------------------------------------

	alldoctorslist := structs.AllDoctorslist_struct{
		Total:   totalDoctors,
		Doctors: Doctors,
	}
	Json_API := structs.AllDoctorslist{
		Message: "Doctors List",
		Status:  true,
		Data:    alldoctorslist,
	}

	jsonData, err := json.Marshal(Json_API)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonData)
}
