package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(groups *models.GroupArr) int {
	rows, err := database.Query("SELECT id, name, about FROM groups")

	if err != nil {
		return 500
	}

	for rows.Next() {
		group := models.Group{}
		rows.Scan(&group.ID, &group.Name, &group.About)
		*groups = append(*groups, &group)
	}
	rows.Close()

	return 200
}

func About(groupName string, groupID int32, group *models.Group) int {
	var err error
	if groupID != 0 {
		err = database.QueryRow("SELECT id, name, about FROM groups WHERE id = $1;", groupID).Scan(&group.ID, &group.Name, &group.About)
	} else {
		err = database.QueryRow("SELECT id, name, about FROM groups WHERE name = $1;", groupName).Scan(&group.ID, &group.Name, &group.About)
	}

	if err == pgx.ErrNoRows {
		return 404
	} else if err != nil {
		return 500
	}

	return 200
}
