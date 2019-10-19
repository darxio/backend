package user

import (
	"backend/internal/database/connection"
	"log"

	"github.com/icrowley/fake"
	"github.com/jackc/pgx"

	"backend/internal/common"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func SignUp(username string, password string) (code int, cookie string, message string) {
	if len(username) < 3 || len(password) < 3 {
		log.Println("database/user.go: 400, Bad username or/and password.")
		return 400, "", "Bad username or/and password."
	}

	var id int
	hashedPassword := common.GeneratePasswordHash(password)
	err := database.QueryRow("INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id;", username, hashedPassword).Scan(&id)
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
	return 201, cookie, "User created successfully."
}

func SignIn(username string, password string) (code int, cookie string, message string) {
	if len(username) < 3 || len(password) < 3 {
		log.Println("database/user.go: 400, Bad username or/and password.")
		return 400, "", "Bad username or/and password."
	}

	var id int
	var scannedPassword string
	err := database.QueryRow("SELECT id, password FROM users WHERE username = $1;", username).Scan(&id, &scannedPassword)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("database/user.go: 404, " + err.Error())
			return 404, "", "User not found."
		}
		log.Println("database/user.go: 500, " + err.Error())
		return 500, "", err.Error()
	}

	if !common.PasswordsMatched(scannedPassword, password) {
		log.Println("database/user.go: 400, Wrong password.")
		return 400, "", "Wrong password."
	}

	cookie = fake.Sentence()
	database.QueryRow("INSERT INTO sessions(user_id, cookie) VALUES ($1, $2);", id, cookie)
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
