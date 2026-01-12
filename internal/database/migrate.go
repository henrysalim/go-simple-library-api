package database

import "database/sql"

func Migrate(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS books (
        id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author VARCHAR(255) NOT NULL,
        year INT NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}
