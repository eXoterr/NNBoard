package server

import "github.com/jmoiron/sqlx"

func SetupDatabase() (*sqlx.DB, error) {
	//Setup database
	db, err := sqlx.Connect("sqlite3", "database.db")

	if err != nil {
		return nil, err
	}

	//Check DB

	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}
