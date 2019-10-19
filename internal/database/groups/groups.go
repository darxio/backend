package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"
	"log"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(groups *models.GroupArr) (code int, message string) {
	rows, err := database.Query(`SELECT id, name, about FROM groups;`)

	if err != nil {
		log.Println("database/groups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		group := models.Group{}
		rows.Scan(&group.ID, &group.Name, &group.About)
		*groups = append(*groups, &group)
	}
	rows.Close()

	return 200, "Successful."
}

func About(groupName string, groupID int32, group *models.Group) (code int, message string) {
	var err error
	if groupID != 0 {
		err = database.QueryRow(`SELECT id, name, about FROM groups WHERE id = $1;`, groupID).Scan(&group.ID, &group.Name, &group.About)
	} else {
		err = database.QueryRow(`SELECT id, name, about FROM groups WHERE name = $1;`, groupName).Scan(&group.ID, &group.Name, &group.About)
	}

	if err == pgx.ErrNoRows {
		log.Println("database/groups.go: 404, " + err.Error())
		return 404, "Group not found."
	} else if err != nil {
		log.Println("database/groups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "Successful."
}
