package queries

// GetUserByEmailAndPasswordQuery get user from user table by email and password
func GetUserByEmailAndPassword() string {
	return "SELECT * FROM users WHERE email = ? AND password = ?"
}

// InsertUser ...
func InsertUser() string {
	return "INSERT INTO users(name, email, password) values(?,?,?)"
}
