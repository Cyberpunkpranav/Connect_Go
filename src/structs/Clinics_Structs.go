package structs

type CLinic_List struct {
	Message string         `json:"message"`
	Status  bool           `json:"status"`
	Data    []Clinics_struct `json:"data"`
}

type Clinics_struct struct {
	Id              *int     `json:"id"`
	Title           *string  `json:"title"`
	Address         *string  `json:"address"`
	Total_roooms    *int     `json:"total_roooms"`
	Phone_number    *string  `json:"phone_number"`
	GSTIN           *string  `json:"gst_IN"`
	State_code      *int     `json:"state_code"`
	Ip_address_list *string  `json:"ip_address_list"`
	Latitude        *float32 `json:"latitude"`
	Longitude       *float32 `json:"longitude"`
	Created_at      *string  `json:"created_at"`
	Updated_at      *string  `json:"updated_at"`
	Created_by      *string  `json:"created_by"`
	Updated_by      *string  `json:"updated_by"`
	Status          *int     `json:"status"`
	Clinic_Rooms []Rooms_struct `json:"clinic_rooms"`
}
type Rooms_struct struct {
	Id          *int    `json:"id"`
	Clinic_id   *int    `json:"clinic_id"`
	Room_number *int    `json:"room_number"`
	Room_type   *int    `json:"room_type "`
	Map_image   *string `json:"map_image "`
	Created_at  *string `json:"created_at"`
	Updated_at  *string `json:"updated_at"`
	Created_by  *string `json:"created_by"`
	Updated_by  *string `json:"updated_by"`
	Status      *int    `json:"status"`
}
