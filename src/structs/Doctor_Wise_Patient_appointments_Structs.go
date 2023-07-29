package structs

type Appointments_Count struct {
	Message string                    `json:"message"`
	Status  bool                      `json:"status"`
	Data    PatientWiseDoctors_struct `json:"data"`
}
type PatientWiseDoctors_struct struct {
	Patient []Patient `json:"patient"`
}
type Patient struct {
	Patient_id   *int    `json:"patient_id"`
	Full_name    *string `json:"patient_name"`
	Phone_number *string `json:"phone_number"`
	Doctor       []Doc   `json:"doctor"`
}
type Doc struct {
	Doctor_id    *int    `json:"doctor_id"`
	Doctor_name  *string `json:"doctor_name"`
	Appointments []Appointment
}
type Appointment struct {
	Appointment_id   *int    `json:"appointment_id"`
	Appointment_date *string `json:"appointment_date"`
}
