package user

import (
	"backend/internal/database/connection"
	"log"

	"github.com/icrowley/fake"
	"github.com/jackc/pgx"

	"backend/internal/common"
	"backend/internal/models"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func SignUp(u *models.User) (code int, cookie string, message string) {
	if len(u.Username) < 3 || len(u.Password) < 3 {
		log.Println("database/user.go: 400, Bad username or/and password.")
		return 400, "", "Bad username or/and password."
	}

	var id int
	hashedPassword := common.GeneratePasswordHash(u.Password)
	err := database.QueryRow("INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id, username;", u.Username, hashedPassword).Scan(&u.ID, &u.Username)
	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			log.Println("database/user.go: 409, " + err.Error())
			return 409, "", "This username already exists."
		}
		log.Println("database/user.go: 500, " + err.Error())
		return 500, "", err.Error()
	}

	cookie = fake.Sentence()
	database.QueryRow("INSERT INTO sessions(user_id, cookie) VALUES ($1, $2);", id, cookie)
	u.Password = ""
	return 201, cookie, "User created successfully."
}

func SignIn(u *models.User) (code int, cookie string, message string) {
	if len(u.Username) < 3 || len(u.Password) < 3 {
		log.Println("database/user.go: 400, Bad username or/and password.")
		return 400, "", "Bad username or/and password."
	}

	var id int
	var scannedPassword string
	err := database.QueryRow("SELECT id, username, password FROM users WHERE username = $1;", u.Username).Scan(&u.ID, &u.Username, &scannedPassword)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("database/user.go: 404, " + err.Error())
			return 404, "", "User not found."
		}
		log.Println("database/user.go: 500, " + err.Error())
		return 500, "", err.Error()
	}

	if !common.PasswordsMatched(scannedPassword, u.Password) {
		log.Println("database/user.go: 400, Wrong password.")
		return 400, "", "Wrong password."
	}

	cookie = fake.Sentence()
	database.QueryRow("INSERT INTO sessions(user_id, cookie) VALUES ($1, $2);", id, cookie)
	u.Password = ""
	return 200, cookie, "User signed in successfully."
}

func SignOut(cookie string) (code int, message string) {
	_, err := database.Exec("DELETE FROM sessions WHERE cookie = $1;", cookie)

	if err != nil {
		log.Println("database/user.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "User signed out successfully."
}
