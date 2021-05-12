package repository

import (
	"log"

	"github.com/OpitiCalvin/novelsTradeIn/pkg/api"
)

// import "github.com/jinzhu/gorm"

// // User type to capture schema and fields for user model
// type User struct {
// 	gorm.Model
// 	Email    string `json:"email"`
// 	UserName string `json:"username"`
// 	password []byte
// }

func (s *storage) CreateUser(request api.NewUserRequest) error {
	newUserStatement := `
		INSERT INTO "user" (username, email, password, fname, surname)
		VALUES ($1, $2, $3, $4, $5);
		`

	err := s.db.QueryRow(newUserStatement, request.Username, request.Email, request.Password, request.FirstName, request.Surname).Err()

	if err != nil {
		log.Printf("create user - there was an error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) GetUsers() ([]api.UserRequest, error) {
	recordsPerPage := 20
	usersStatement := `
		SELECT id, username, email, fname, surname
		FROM "user"
		LIMIT $1
	`

	var users []api.UserRequest
	rows, err := s.db.Query(usersStatement, recordsPerPage)

	if err != nil {
		log.Printf("get users - there was an error: %v", err.Error())
		return []api.UserRequest{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id                              int
			username, email, fname, surname string
		)

		err := rows.Scan(&id, &username, &email, &fname, &surname)
		if err != nil {
			log.Printf("Get Users - Error while scanning result set rows: %v", err.Error())
			return []api.UserRequest{}, err
		}

		users = append(users, api.UserRequest{ID: id, Username: username, Email: email, FirstName: fname, Surname: surname})
	}

	if err := rows.Err(); err != nil {
		log.Printf("get users - error with rows: %v", err.Error())
		return []api.UserRequest{}, err
	}

	return users, nil
}

func (s *storage) GetUser(userID int) (api.UserRequest, error) {
	getUserStatement := `
	SELECT id, username, email, fname, surname
	FROM user
	WHERE id=$1
	`

	var user api.UserRequest

	err := s.db.QueryRow(getUserStatement, userID).Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.Username)

	if err != nil {
		log.Printf("get user - this was the error: %v", err.Error())
		return api.UserRequest{}, err
	}

	return user, nil
}
