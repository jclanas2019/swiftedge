package helpers

import (
	"database/sql"
	"fmt"
)

// ConnectDB establece una conexión a la base de datos MySQL.
func ConnectDB(user, password, host, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error verificando la conexión: %v", err)
	}
	return db, nil
}
