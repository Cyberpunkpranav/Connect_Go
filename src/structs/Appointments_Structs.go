package structs

// Appointments/List
type Appointment_List struct {
	Message string                `json:"message"`
	Status  bool                  `json:"status"`
	Data    []Appointments_struct `json:"data"`
}

type Appointments_struct struct {
	Id                         *int                      `json:"id"`
	Bill_id                    *int                      `json:"bill_id"`
	Doctor_id                  *int                      `json:"doctor_id"`
	Patient_id                 *int                      `json:"patient_id"`
	Timeslot_id                *int                      `json:"timeslot_id"`
	Clinic_id                  *int                      `json:"clinic_id"`
	Payment_method             *string                   `json:"payment_method"`
	Payment_method_details     *string                   `json:"payment_method_details"`
	Pay_id                     *string                   `json:"pay_id"`
	Appointment_status         *int                      `json:"appointment_status"`
	Payment_status             *int                      `json:"payment_status"`
	Cons_fee                   *float32                  `json:"cons_fee"`
	Show_cons_fee              *string                   `json:"show_cons_fee"`
	Cons_text                  *string                   `json:"cons_text"`
	Doc_discount               *float32                  `json:"doc_discount"`
	Discount                   *float32                  `json:"discount"`
	Coupon_id                  *int                      `json:"coupon_id"`
	Aartas_discount            *float32                  `json:"aartas_discount"`
	Appointment_type           *int                      `json:"appointment_type"`
	Camp_id                    *int                      `json:"camp_id"`
	Camp_doctor_id             *int                      `json:"camp_doctor_id"`
	Procedure_id               *int                      `json:"procedure_id"`
	Procedure_cost             *string                   `json:"procedure_cost"`
	Patient_bundles_id         *int                      `json:"patient_bundles_id"`
	Appointment_date           *string                   `json:"appointment_date"`
	SGST                       *float32                  `json:"sgst"`
	CGST                       *float32                  `json:"cgst"`
	Total_amount               *float32                  `json:"total_amount"`
	Patient_rewards_history_id *int                      `json:"patient_rewards_history_id"`
	Whatsapp_sent              *int                      `json:"Whatsapp_sent"`
	Is_confirmed               *int                      `json:"is_confirmed"`
	Prescription_file          *string                   `json:"prescription_file"`
	Bill_file                  *string                   `json:"bill_file"`
	Cons_start_time            *string                   `json:"cons_start_time"`
	Cons_end_time              *string                   `json:"cons_end_time"`
	Appointment_visit_type     *string                   `json:"appointment_visit_type"`
	Notes                      *string                   `json:"notes"`
	Nursing_notes              *string                   `json:"nursing_notes"`
	Follow_up_date             *string                   `json:"follow_up_date"`
	Created_at                 *string                   `json:"created_at"`
	Updated_at                 *string                   `json:"updated_at"`
	Created_by                 *string                   `json:"created_by"`
	Updated_by                 *string                   `json:"updated_by"`
	Status                     *int                      `json:"status"`
	Patient                    Patient_struct            `json:"patient"`
	Timeslot                   Timeslot_struct           `json:"timeslot"`
	Doctor                     Doctors_struct            `json:"doctor"`
	OtherCharges               []Other_charges_struct    `json:"other_charges"`
	PendingPayments            []Pending_payments_struct `json:"pending_payments"`
}

type Patient_struct struct {
	Id                 *int    `json:"id"`
	Full_name          *string `json:"full_name"`
	Phone_country_code *string `json:"phone_country_code"`
	Phone_number       *string `json:"phone_number"`
	Email              *string `json:"email"`
	Gender             *string `json:"gender"`
	Pin_code           *string `json:"pin_code"`
	Fcm_token          *string `json:"fcm_token"`
	Device_id          *string `json:"device_id"`
	Version            *string `json:"version"`
	Platform           *string `json:"platform"`
	Dob                *string `json:"dob"`
	Age                *int    `json:"age"`
	Relation           *string `json:"relation"`
	Link_id            *string `json:"link_id"`
	Location           *string `json:"location"`
	Latitude           *string `json:"latitude"`
	Longitude          *string `json:"longitude"`
	Membership_type_id *string `json:"membership+type_id"`
	Source             *string `json:"source"`
	Created_at         *string `json:"created_at"`
	Updated_at         *string `json:"updated_at"`
	Created_by         *string `json:"created_by"`
	Updated_by         *string `json:"updated_by"`
	Status             *int    `json:"status"`
	// Address            Address_struct `json:"address"`
	// Check_in_details   *string          `json:"check_in_details"`
	// Reward_points      *int             `json:"reward_points"`
}
type Timeslot_struct struct {
	Id              *int          `json:"id"`
	Doctor_id       *int          `json:"doctor_id"`
	Date            *string       `json:"date"`
	Time_from       *string       `json:"time_from"`
	Time_to         *string       `json:"time_to"`
	Clinic_id       *int          `json:"clinic_id"`
	Clinic_rooms_id *int          `json:"clinic_rooms_id"`
	Booking_status  *int          `json:"booking_status"`
	Created_at      *string       `json:"created_at"`
	Updated_at      *string       `json:"updated_at"`
	Created_by      *string       `json:"created_by"`
	Updated_by      *string       `json:"updated_by"`
	Status          *int          `json:"status"`
	Clinic          Clinic_struct `json:"clinic"`
}

type Doctors_struct struct {
	Id                       *string `json:"id"`
	Doctor_name              *string `json:"doctor_name"`
	Phone_number             *string `json:"phone_number"`
	Gender                   *string `json:"gender"`
	Address                  *string `json:"address"`
	Languages                *string `json:"languages"`
	Degree_suffix            *string `json:"degree_suffix"`
	Expertise_subSpecialty   *string `json:"expertise_subSpecialty"`
	ConsulationFee           *string `json:"consulationFee"`
	Appointment_phone_number *string `json:"appointment_phone_number"`
	Status                   *string `json:"status"`
}
type Clinic_struct struct {
	Id              *int    `json:"id"`
	Title           *string `json:"title"`
	Address         *string `json:"address"`
	Total_rooms     *int    `json:"total_rooms"`
	Phone_number    *int    `json:"phone_number"`
	GSTIN           *string `json:"GSTIN"`
	State_code      *string `json:"State_code"`
	Ip_address_list *string `json:"ip_address_list"`
	Latitude        *string `json:"latitude"`
	Longitude       *string `json:"longitude"`
	Created_at      *string `json:"created_at"`
	Updated_at      *string `json:"updated_at"`
	Created_by      *string `json:"created_by"`
	Updated_by      *string `json:"updated_by"`
	Status          *int    `json:"status"`
}
type Other_charges_struct struct {
	Id             int      `json:"id"`
	Appointment_id *int     `json:"appointment_id"`
	Description    *string  `json:"description		"`
	Total_amount   *float32 `json:"	total_amount"`
	Discount       *float32 `json:"	discount"`
	Amount         *float32 `json:"	amount"`
	Gst_rate       *float32 `json:"	gst_rate"`
	Final_amount   *float32 `json:"	final_amount"`
	Created_at     *string  `json:"	created_at"`
	Updated_at     *string  `json:"	updated_at"`
	Created_by     *string  `json:"	created_by"`
	Updated_by     *string  `json:"	updated_by"`
	Status         *string  `json:"	status"`
}
type Pending_payments_struct struct {
	Id                     *int     `json:"id"`
	Appointment_id         *int     `json:"appointment_id"`
	Pending_amount         *float32 `json:"pending_amount"`
	Pending_date           *string  `json:"pending_date"`
	Paid_amount            *float32 `json:"paid_amount"`
	Paid_date              *string  `json:"paid_date"`
	Is_Paid                *int     `json:"Is_Paid"`
	Payment_method_details *string  `json:"Payment_method_details"`
	Created_at             *string  `json:"	created_at"`
	Updated_at             *string  `json:"	updated_at"`
	Created_by             *string  `json:"	created_by"`
	Updated_by             *string  `json:"	updated_by"`
	Status                 *int     `json:"	status"`
}

type Payment_Method_Details struct {
	Cash     float32 `json:"cash"`
	Card     float32 `json:"card"`
	RazorPay float32 `json:"razorpay"`
	Paytm    float32 `json:"paytm"`
	Phonepay float32 `json:"phonepay"`
}
