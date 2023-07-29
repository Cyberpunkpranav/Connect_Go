package structs

type Patient_List struct {
	Message string       `json:"message"`
	Status  bool         `json:"status"`
	Data    Patient_Data `json:"patients_list"`
}
type Patient_Data struct {
	Total_Count int           `json:"total_count"`
	Data        []AllPatients `json:"data"`
}
type AllPatients struct {
	Id                 *int             `json:"id"`
	Full_name          *string          `json:"full_name"`
	Phone_country_code *int             `json:"phone_country_code"`
	Phone_number       *string          `json:"phone_number"`
	Email              *string          `json:"email"`
	Gender             *string          `json:"gender"`
	Pin_code           *int             `json:"pin_code"`
	Fcm_token          *string          `json:"fcm_token"`
	Device_id          *string          `json:"device_id"`
	Version            *float32         `json:"version"`
	Platform           *string          `json:"platform"`
	Dob                *string          `json:"dob"`
	Age                *int             `json:"age"`
	Relation           *string          `json:"relation"`
	Link_id            *string          `json:"link_id"`
	Location           *string          `json:"location"`
	Latitude           *float64         `json:"latitude"`
	Longitude          *float64         `json:"longitude"`
	Membership_type_id *int             `json:"membership_type_id"`
	Source             *string          `json:"source"`
	Created_at         *string          `json:"created_at"`
	Updated_at         *string          `json:"updated_at"`
	Created_by         *string          `json:"created_by"`
	Updated_by         *string          `json:"updated_by"`
	Status             *int             `json:"status"`
	Address            []Address_struct `json:"address"`
}
type Address_struct struct {
	Id            *int    `json:"id"`
	Patient_id    *int    `json:"patient_id"`
	Full_name     *string `json:"full_name"`
	Address_line1 *string `json:"address_line1"`
	Address_line2 *string `json:"address_line2"`
	Zip_code      *string `json:"zip_code"`
	Country       *string `json:"country"`
	State         *string `json:"state"`
	City          *string `json:"city"`
	Created_at    *string `json:"created_at"`
	Updated_at    *string `json:"updated_at"`
	Status        *int    `json:"status"`
}
