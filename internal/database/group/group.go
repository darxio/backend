package group

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
	"log"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func About(groupName string, group *models.Group) int {
	err := database.QueryRow("SELECT name, about FROM groups WHERE name = $1;", groupName).Scan(&group.Name, &group.About)

	log.Println(err)
	if err == pgx.ErrNoRows {
		return 404
	} else if err != nil {
		return 500
	}

	return 200
}
