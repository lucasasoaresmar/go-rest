package database

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/lucasasoaresmar/go-rest/adapters/database/queries"
	"github.com/lucasasoaresmar/go-rest/models"
)

// User implements UserRepository
type User struct{}

// FindByEmailAndPassword ...
func (u *User) FindByEmailAndPassword(email string, password string) (user *models.User, err error) {
	execute(func(db *sql.DB) {
		var id int
		user = new(models.User)

		// rows, err := db.Query("SELECT email, password FROM users")
		// var temail string
		// var tpass string

		// for rows.Next() {
		// 	err := rows.Scan(&temail, &tpass)
		// 	if err != nil {
		// 		fmt.Print(err)
		// 		return
		// 	}

		// 	fmt.Println(temail, tpass, email, password)
		// }

		err = db.QueryRow(queries.GetUserByEmailAndPassword(), email, password).
			Scan(&id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			fmt.Print(err)
			return
		}

		user.ID = strconv.Itoa(id)
	})

	return user, err
}
