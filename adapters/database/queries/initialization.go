package queries

// InitializeDB ...
func InitializeDB() string {
	return `
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			name VARCHAR(255),
			email VARCHAR(255),
			password VARCHAR(255)
    );`
}
