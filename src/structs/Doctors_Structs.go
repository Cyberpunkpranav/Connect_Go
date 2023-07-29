package structs

type AllDoctorslist struct {
	Message string                `json:"message"`
	Status  bool                  `json:"status"`
	Data    AllDoctorslist_struct `json:"data"`
}

type AllDoctorslist_struct struct {
	Total   int           `json:"total_count"`
	Doctors []Doctor_info `json:"doctors_list"`
}

type Doctor_info struct {
	Id                       *int    `json:"id"`
	User_id                  *string `json:"user_id"`
	Clinic_id                *int    `json:"clinic_id"`
	Room_id                  *int    `json:"room_id"`
	Doctor_name              *string `json:"doctor_name"`
	Email                    *string `json:"email"`
	Phone_country_code       *string `json:"phone_country_code"`
	Phone_number             *string `json:"phone_number"`
	Password                 *string `json:"password"`
	Fcm_token                *string `json:"fcm_token"`
	Gender                   *string `json:"gender"`
	Address                  *string `json:"address"`
	City_id                  *int    `json:"city_id"`
	Languages                *string `json:"languages"`
	Speciality_id            *int    `json:"speciality_id"`
	Registration_num         *string `json:"registration_num"`
	Medical_authority_id     *int    `json:"medical_authority_id"`
	Degree_suffix            *string `json:"degree_suffix"`
	Expertise_subSpecialty   *string `json:"expertise_subSpecialty"`
	Image                    *string `json:"image"`
	Pdf_header               *string `json:"pdf_header"`
	Pdf_footer               *string `json:"pdf_footer"`
	ScheduleDays             *string `json:"scheduleDays"`
	ScheduleTimings          *string `json:"scheduleTimings"`
	Hours_per_month          *string `json:"hours_per_month"`
	Agreement_file           *string `json:"agreement_file"`
	Start_date               *string `json:"start_date"`
	Agreement_renew_date     *string `json:"agreement_renew_date"`
	ConsulationFee           *string `json:"consulationFee"`
	Time_per_patient         *string `json:"time_per_patient"`
	Appointment_phone_number *string `json:"appointment_phone_number"`
	Website                  *string `json:"website"`
	Instagram                *string `json:"instagram"`
	Facebook                 *string `json:"facebook"`
	Youtube                  *string `json:"youtube"`
	Profile_share_count      *string `json:"profile_share_count"`
	Profile_verified         *int    `json:"profile_verified"`
	Profile_status           *int    `json:"profile_status"`
	Reset_token              *string `json:"reset_token"`
	On_request               *string `json:"on_request"`
	Pre_pay_enrolled         *string `json:"pre_pay_enrolled"`
	Pre_pay_discount         *string `json:"pre_pay_discount"`
	Total_experience         *string `json:"total_experience"`
	About                    *string `json:"about"`
	Created_at               *string `json:"created_at"`
	Updated_at               *string `json:"updated_at"`
	Created_by               *string `json:"created_by"`
	Updated_by               *string `json:"updated_by"`
	Status                   *int    `json:"status"`
	Rent_amount              *string `json:"rent_amount"`
	Security_amount          *string `json:"security_amount"`
	Speciality               Speciality
	Timings                  []Timings_struct `json:"timings"`
}

type Speciality struct {
	Id          *int    `json:"id"`
	Name        *string `json:"name"`
	Icon_url    *string `json:"icon_url"`
	Status      *int    `json:"status"`
	Description *string `json:"description"`
	Conditions  *string `json:"conditions"`
	Created_at  *string `json:"created_at"`
	Updated_at  *string `json:"updated_at"`
	Bg_color1   *string `json:"bg_color1"`
	Bg_color2   *string `json:"bg_color2"`
	Icon_color  *string `json:"icon_color"`
	Created_by  *string `json:"created_by"`
	Updated_by  *string `json:"updated_by"`
}
type Timings_struct struct {
	Id               *int          `json:"id              "`
	Doctor_id        *int          `json:"doctor_id       "`
	ScheduleDays     *string       `json:"scheduleDays    "`
	ScheduleTimings  *string       `json:"scheduleTimings "`
	Time_per_patient *string       `json:"time_per_patient"`
	Clinic_id        *int          `json:"clinic_id       "`
	Clinic_rooms_id  *int          `json:"clinic_rooms_id "`
	Created_by       *string       `json:"created_by      "`
	Updated_by       *string       `json:"updated_by      "`
	Created_at       *string       `json:"created_at      "`
	Updated_at       *string       `json:"updated_at      "`
	Status           *int          `json:"status          "`
	Clinic           Clinic_struct `json:"clinic"`
}
