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
		return 400, "", "Bad username or/and password."
	}

	var id int
	hashedPassword := common.GeneratePasswordHash(password)
	err := database.QueryRow("INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id;", username, hashedPassword).Scan(&id)
	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			return 409, "", "This username already exists."
		}
		return 500, "", "Something went wrong.."
	}

	cookie = fake.Sentence()
	database.QueryRow("INSERT INTO sessions(user_id, cookie) VALUES ($1, $2);", id, cookie)
	return 201, cookie, "User created successfully."
}

func SignIn(username string, password string) (code int, cookie string, message string) {
	if len(username) < 3 || len(password) < 3 {
		return 400, "", "Bad username or/and password."
	}

	var id int
	var scannedPassword string
	err := database.QueryRow("SELECT id, password FROM users WHERE username = $1;", username).Scan(&id, &scannedPassword)

	if err == pgx.ErrNoRows {
		return 404, "", "User not found."
	} else {
		if !common.PasswordsMatched(scannedPassword, password) {
			return 400, "", "Wrong password."
		}
		cookie = fake.Sentence()
		database.QueryRow("INSERT INTO sessions(user_id, cookie) VALUES ($1, $2);", id, cookie)
		return 200, cookie, "User signed in successfully."
	}

	return 500, "", "Something went wrong.."
}

func SignOut(cookie string) (code int, message string) {
	log.Println(cookie)
	_, err := database.Exec("DELETE FROM sessions WHERE cookie = $1;", cookie)

	if err != nil {
		return 500, "Something went wrong.."
	}

	return 200, "User signed out successfully."
}