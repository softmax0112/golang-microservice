package user

import "database/sql"

//DBConfig struct holds the SQL Driver info
type DBConfig struct {
	DB *sql.DB
}

// func (d UserConfig) CreateUser(username string, password string) error {

// }
