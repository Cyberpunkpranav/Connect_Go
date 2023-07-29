package controllers

import (
	sql "ConnectApp/src/sql"
	structs "ConnectApp/src/structs"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Appointments_List(w http.ResponseWriter, r *http.Request) {
	var query string
	var params []interface{}
	//----------------------------------Database Query---------------------------------------------------
	db := sql.DB
	var (
		doctor_id          = r.URL.Query().Get("doctor_id")
		from_date          = r.URL.Query().Get("from_date")
		to_date            = r.URL.Query().Get("to_date")
		clinic_id          = r.URL.Query().Get("clinic_id")
		appointment_status = r.URL.Query().Get("appointment_status")
	)
	if from_date == "" {
		from_date = time.Now().Format("2006-01-02")
		to_date = time.Now().Format("2006-01-02")
	}
	if to_date == "" {
		to_date = r.URL.Query().Get("from_date")
	}

	if doctor_id != "" && clinic_id != "" && from_date != "" && to_date != "" && appointment_status != "" {
		query = `
		SELECT appointments.* ,patients.*,doctor_timeslots.*,clinic.*, doctor.id, doctor.doctor_name,doctor.phone_number,doctor.gender,doctor.address,doctor.languages,doctor.degree_suffix,doctor.expertise_subSpecialty,doctor.consulationFee,doctor.appointment_phone_number,doctor.status FROM appointments
join patients on appointments.patient_id = patients.id
join doctor_timeslots on appointments.timeslot_id =doctor_timeslots.id
join clinic on appointments.clinic_id = clinic.id
join doctor on appointments.doctor_id = doctor.id
WHERE appointments.appointment_date BETWEEN ? AND ? AND appointments.doctor_id = ? AND appointments.clinic_id = ? AND appointments.appointment_status = ? ;
 `
		params = []interface{}{from_date, to_date, doctor_id, clinic_id, appointment_status}
	} else if doctor_id != "" && from_date != "" && to_date != "" && clinic_id != "" && appointment_status == "" {
		query = `
				 SELECT appointments.* ,patients.*,doctor_timeslots.*,clinic.*, doctor.id, doctor.doctor_name,doctor.phone_number,doctor.gender,doctor.address,doctor.languages,doctor.degree_suffix,doctor.expertise_subSpecialty,doctor.consulationFee,doctor.appointment_phone_number,doctor.status  FROM appointments
				 JOIN patients ON appointments.patient_id = patients.id 
				 join doctor_timeslots on appointments.timeslot_id = doctor_timeslots.id 
				 join clinic on appointments.clinic_id = clinic.id 
				 join doctor on appointments.doctor_id = doctor.id
				 WHERE appointments.appointment_date BETWEEN ? AND ? AND appointments.doctor_id = ? AND appointments.clinic_id = ? `
		params = []interface{}{from_date, to_date, doctor_id, clinic_id}
	} else if doctor_id != "" && from_date != "" && to_date != "" && clinic_id == "" && appointment_status == "" {
		query = `
		SELECT appointments.* ,patients.*,doctor_timeslots.*,clinic.*, doctor.id, doctor.doctor_name,doctor.phone_number,doctor.gender,doctor.address,doctor.languages,doctor.degree_suffix,doctor.expertise_subSpecialty,doctor.consulationFee,doctor.appointment_phone_number,doctor.status  FROM appointments
				 JOIN patients ON appointments.patient_id = patients.id
				 join doctor_timeslots on appointments.timeslot_id =doctor_timeslots.id
				 join clinic on appointments.clinic_id = clinic.id
				 join doctor on appointments.doctor_id = doctor.id
				 WHERE appointments.appointment_date BETWEEN ? AND ? AND appointments.doctor_id = ? `
		params = []interface{}{from_date, to_date, doctor_id}
	} else if doctor_id == "" && from_date != "" && to_date != "" && clinic_id == "" && appointment_status == "" {
		query = `
		SELECT appointments.* ,patients.*,doctor_timeslots.*,clinic.*, doctor.id, doctor.doctor_name,doctor.phone_number,doctor.gender,doctor.address,doctor.languages,doctor.degree_suffix,doctor.expertise_subSpecialty,doctor.consulationFee,doctor.appointment_phone_number,doctor.status  FROM appointments
		JOIN patients ON appointments.patient_id = patients.id
		join doctor_timeslots on appointments.timeslot_id =doctor_timeslots.id
		join clinic on appointments.clinic_id = clinic.id
		join doctor on appointments.doctor_id = doctor.id
		WHERE appointments.appointment_date BETWEEN ? AND ?  `
		params = []interface{}{from_date, to_date}
	} else {
		query = `
		SELECT appointments.* ,patients.*,doctor_timeslots.*,clinic.*, doctor.id, doctor.doctor_name,doctor.phone_number,doctor.gender,doctor.address,doctor.languages,doctor.degree_suffix,doctor.expertise_subSpecialty,doctor.consulationFee,doctor.appointment_phone_number,doctor.status  FROM appointments
		JOIN patients ON appointments.patient_id = patients.id
		join doctor_timeslots on appointments.timeslot_id =doctor_timeslots.id
		join clinic on appointments.clinic_id = clinic.id
		join doctor on appointments.doctor_id = doctor.id
		 WHERE appointments.appointment_date BETWEEN ? AND ?  `
		params = []interface{}{from_date, to_date}
	}

	dataDB, err := db.Query(query, params...)
	if err != nil {
		log.Fatal(err)
	}
	//----------------------------------Database Query---------------------------------------------------

	//---------------------------------Getting Appointments-----------------------------------------------

	var Appointments []structs.Appointments_struct
	// var Othercharges []Other_charges_struct
	for dataDB.Next() {
		var appointment structs.Appointments_struct
		var othercharge structs.Other_charges_struct
		var pendingpayment structs.Pending_payments_struct
		err := dataDB.Scan(
			&appointment.Id,
			&appointment.Bill_id,
			&appointment.Doctor_id,
			&appointment.Patient_id,
			&appointment.Timeslot_id,
			&appointment.Clinic_id,
			&appointment.Payment_method,
			&appointment.Payment_method_details,
			&appointment.Pay_id,
			&appointment.Appointment_status,
			&appointment.Payment_status,
			&appointment.Cons_fee,
			&appointment.Show_cons_fee,
			&appointment.Cons_text,
			&appointment.Doc_discount,
			&appointment.Discount,
			&appointment.Coupon_id,
			&appointment.Aartas_discount,
			&appointment.Appointment_type,
			&appointment.Camp_id,
			&appointment.Camp_doctor_id,
			&appointment.Procedure_id,
			&appointment.Procedure_cost,
			&appointment.Patient_bundles_id,
			&appointment.Appointment_date,
			&appointment.SGST,
			&appointment.CGST,
			&appointment.Total_amount,
			&appointment.Patient_rewards_history_id,
			&appointment.Whatsapp_sent,
			&appointment.Is_confirmed,
			&appointment.Prescription_file,
			&appointment.Bill_file,
			&appointment.Cons_start_time,
			&appointment.Cons_end_time,
			&appointment.Appointment_visit_type,
			&appointment.Notes,
			&appointment.Nursing_notes,
			&appointment.Follow_up_date,
			&appointment.Created_at,
			&appointment.Updated_at,
			&appointment.Created_by,
			&appointment.Updated_by,
			&appointment.Status,
			&appointment.Patient.Id,
			&appointment.Patient.Full_name,
			&appointment.Patient.Phone_country_code,
			&appointment.Patient.Phone_number,
			&appointment.Patient.Email,
			&appointment.Patient.Gender,
			&appointment.Patient.Pin_code,
			&appointment.Patient.Fcm_token,
			&appointment.Patient.Device_id,
			&appointment.Patient.Version,
			&appointment.Patient.Platform,
			&appointment.Patient.Dob,
			&appointment.Patient.Age,
			&appointment.Patient.Relation,
			&appointment.Patient.Link_id,
			&appointment.Patient.Location,
			&appointment.Patient.Latitude,
			&appointment.Patient.Longitude,
			&appointment.Patient.Membership_type_id,
			&appointment.Patient.Source,
			&appointment.Patient.Created_at,
			&appointment.Patient.Updated_at,
			&appointment.Patient.Created_by,
			&appointment.Patient.Updated_by,
			&appointment.Patient.Status,
			&appointment.Timeslot.Id,
			&appointment.Timeslot.Doctor_id,
			&appointment.Timeslot.Date,
			&appointment.Timeslot.Time_from,
			&appointment.Timeslot.Time_to,
			&appointment.Timeslot.Clinic_id,
			&appointment.Timeslot.Clinic_rooms_id,
			&appointment.Timeslot.Booking_status,
			&appointment.Timeslot.Created_at,
			&appointment.Timeslot.Updated_at,
			&appointment.Timeslot.Created_by,
			&appointment.Timeslot.Updated_by,
			&appointment.Timeslot.Status,
			&appointment.Timeslot.Clinic.Id,
			&appointment.Timeslot.Clinic.Title,
			&appointment.Timeslot.Clinic.Address,
			&appointment.Timeslot.Clinic.Total_rooms,
			&appointment.Timeslot.Clinic.Phone_number,
			&appointment.Timeslot.Clinic.GSTIN,
			&appointment.Timeslot.Clinic.State_code,
			&appointment.Timeslot.Clinic.Ip_address_list,
			&appointment.Timeslot.Clinic.Latitude,
			&appointment.Timeslot.Clinic.Longitude,
			&appointment.Timeslot.Clinic.Created_at,
			&appointment.Timeslot.Clinic.Updated_at,
			&appointment.Timeslot.Clinic.Created_by,
			&appointment.Timeslot.Clinic.Updated_by,
			&appointment.Timeslot.Clinic.Status,
			&appointment.Doctor.Id,
			&appointment.Doctor.Doctor_name,
			&appointment.Doctor.Phone_number,
			&appointment.Doctor.Gender,
			&appointment.Doctor.Address,
			&appointment.Doctor.Languages,
			&appointment.Doctor.Degree_suffix,
			&appointment.Doctor.Expertise_subSpecialty,
			&appointment.Doctor.ConsulationFee,
			&appointment.Doctor.Appointment_phone_number,
			&appointment.Doctor.Status,
		)
		if err != nil {
			panic(err.Error())
		}

		otherchargeDB, err := db.Query(`SELECT * FROM appointments_other_charges WHERE appointment_id = ? `, *appointment.Id)
		if err != nil {
			panic(err.Error())
		}
		for otherchargeDB.Next() {
			err := otherchargeDB.Scan(
				&othercharge.Id,
				&othercharge.Appointment_id,
				&othercharge.Description,
				&othercharge.Total_amount,
				&othercharge.Discount,
				&othercharge.Amount,
				&othercharge.Gst_rate,
				&othercharge.Final_amount,
				&othercharge.Created_at,
				&othercharge.Updated_at,
				&othercharge.Created_by,
				&othercharge.Updated_by,
				&othercharge.Status,
			)
			if err != nil {
				log.Fatal(err)
			}
			appointment.OtherCharges = append(appointment.OtherCharges, othercharge)
		}
		otherchargeDB.Close()
		pendingpaymentDB, err := db.Query(`SELECT * FROM appointments_pending_payments WHERE appointment_id = ?  `, *appointment.Id)
		if err != nil {
			log.Fatal(err)
		}
		for pendingpaymentDB.Next() {
			err := pendingpaymentDB.Scan(
				&pendingpayment.Id,
				&pendingpayment.Appointment_id,
				&pendingpayment.Pending_amount,
				&pendingpayment.Pending_date,
				&pendingpayment.Paid_amount,
				&pendingpayment.Paid_date,
				&pendingpayment.Is_Paid,
				&pendingpayment.Payment_method_details,
				&pendingpayment.Created_at,
				&pendingpayment.Updated_at,
				&pendingpayment.Created_by,
				&pendingpayment.Updated_by,
				&pendingpayment.Status,
			)
			if err != nil {
				panic(err.Error())
			}

			appointment.PendingPayments = append(appointment.PendingPayments, pendingpayment)
		}
		pendingpaymentDB.Close()
		if err := dataDB.Err(); err != nil {
			panic(err.Error())
		}
		Appointments = append(Appointments, appointment)

	}
	//---------------------------------Getting Appointments-----------------------------------------------

	APIData := structs.Appointment_List{
		Message: "Appointments List",
		Status:  true,
		Data:    Appointments,
	}

	jsonData, err := json.Marshal(APIData)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	db.Close()
	w.Write(jsonData)

}
