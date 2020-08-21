package user

import "database/sql"

//DBConfig struct holds the SQL Driver info
type DBConfig struct {
	DB *sql.DB
}
