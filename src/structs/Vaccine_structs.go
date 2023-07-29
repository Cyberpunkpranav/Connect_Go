package structs

//Vaccines/List
type Vaccines_struct struct {
	Id                  int      `json:"id"`
	Vaccine_name        string   `json:"name"`
	Vaccine_category_id int      `json:"vaccine_category_id"`
	Vaccine_stocks      []string `json:"vaccine_stocks"`
	Created_at          string   `json:"created_at"`
	Updated_at          string   `json:"updated_at"`
	Created_by          string   `json:"created_by"`
	Updated_by          string   `json:"updated_by"`
	Status              int      `json:"status"`
}
type Vaccine_Stocks struct {
	Stock_id          int     `json:"id"`
	Vaccine_brand_id  int     `json:"vaccine_brand_id"`
	Purchase_entry_id int     `json:"purchase_entry_id"`
	Qty               int     `json:"qty"`
	Free_qty          int     `json:"free_qty"`
	Rate              float64 `json:"rate"`
	Expiry_date       string  `json:"expiry_date"`
	Batch_no          string  `json:"batch_no"`
	Mfd_date          string  `json:"mfd_date"`
	Mrp               float32 `json:"mrp"`
	Cost              float32 `json:"cost"`
	Sgst              float32 `json:"SGST"`
	Sgst_rate         float32 `json:"SGST_rate"`
	Cgst              float32 `json:"CGST"`
	Cgst_rate         float32 `json:"CGST_rate"`
	Igst              float32 `json:"IGST"`
	Igst_rate         float32 `json:"IGST_rate"`
	Discount          float32 `json:"discount"`
	Location_id       int     `json:"location_id"`
	Channel           int     `json:"channel"`
	Barcode_no        string  `json:"barcode_no"`
	Trade_discount    float32 `json:"trade_discount"`
	Current_stock     int     `json:"current_stock"`
	Total_amount      float32 `json:"total_amount"`
	Created_at        string  `json:"created_at"`
	Updated_at        string  `json:"updated_at"`
	Created_by        string  `json:"created_by"`
	Updated_by        string  `json:"updated_by"`
}
