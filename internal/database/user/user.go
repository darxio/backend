package user

import (
	"backend/internal/database/connection"
	"log"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func SignUp(username string, password string) (code int, cookie string, message string) {
	log.Println(len(username))
	log.Println(len(password))
	if len(username) < 3 || len(password) < 3 {
		return 400, "", "Bad username or/and password."
	}
	var id int

	err := database.QueryRow("INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id;", username, password).Scan(&id)

	log.Println(err)
	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			return 409, "", "This username already exists."
		}
		return 500, "", "Something went wrong.."
	}

	cookie = "temp_cookie"
	database.QueryRow("INSERT INTO cookies(user_id, cookie) VALUES ($1, $2);", id, cookie)
	return 201, cookie, "User created successfully."
}

func SignIn(username string, password string) (code int, cookie string, message string) {
	if len(username) < 3 || len(password) < 3 {
		return 400, "", "Bad username or/and password."
	}

	// kostyl
	id := -1
	database.QueryRow("SELECT id FROM users WHERE username = $1 AND password = $2;", username, password).Scan(&id)

	if id == -1 {
		return 404, "", "User not found."
	} else {
		cookie = "temp_cookie"
		database.QueryRow("INSERT INTO cookies(user_id, cookie) VALUES ($1, $2);", id, cookie)
		return 200, cookie, "User signed in successfully."
	}

	return 500, "", "Something went wrong.."
}

func SignOut(cookie string) (code int, message string) {
	err := database.QueryRow("DELETE FROM sessions WHERE cookie = $1;", cookie)

	if err != nil {
		return 500, "Something went wrong.."
	}

	return 200, "User signed out successfully."
}
