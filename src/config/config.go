package config

type Database struct {
	Username   string
	Database   string
	Password   string
	Protocol   string
	Port       string
	Ip_address string
}

func Db_Config() Database {
	var db Database
	db.Username = "root"
	db.Database = "aartas"
	db.Password = "MYSQLaccount123"
	db.Protocol = "tcp"
	db.Ip_address = "127.0.0.1"
	db.Port = "3306"
	return db
}