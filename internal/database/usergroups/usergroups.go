package groups

import (
	"backend/internal/database/connection"
	"backend/internal/models"
	"log"

	_ "log"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(cookie string, groups *models.GroupArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/usergroups.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	rows, err := database.Query(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE user_id = $1;`, id)

	if err != nil {
		log.Println("database/usergroups.go: 500, " + err.Error())
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

func About(cookie string, groupName string, groupID int32, group *models.Group) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/usergroups.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	var err error
	if groupID != 0 {
		err = database.QueryRow(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE user_id = $1
								AND group_id = $2;`, id, groupID).Scan(&group.ID, &group.Name, &group.About)
	} else {
		err = database.QueryRow(`SELECT id, name, about
								FROM groups
								JOIN user_groups
								ON groups.id = user_groups.group_id
								WHERE user_id = $1
								AND name = $2;`, id, groupName).Scan(&group.ID, &group.Name, &group.About)
	}

	if err == pgx.ErrNoRows {
		log.Println("database/usergroups.go: 404, " + err.Error())
		return 404, "Group not found."
	} else if err != nil {
		log.Println("database/usergroups.go: 500, " + err.Error())
		return 500, "Something went wrong.."
	}

	return 200, "Successful."
}

func Add(cookie string, groupName string, groupID int32, groups *models.GroupArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/usergroups.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	var err error
	if groupID != 0 {
		_, err = database.Exec("INSERT INTO user_groups(user_id, group_id) VALUES ($1, $2);", id, groupID)
	} else {
		_, err = database.Exec("INSERT INTO user_groups(user_id, group_id) VALUES ($1, (SELECT id FROM groups WHERE name = $2));", id, groupName)
	}

	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			log.Println("database/usergroups.go: 409, " + err.Error())
			return 409, "This user is already in this group."
		}
		if pgErr.Code == "23503" {
			log.Println("database/usergroups.go: 404, " + err.Error())
			return 404, "This group doesn't exist."
		}
		log.Println("database/usergroups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return All(cookie, groups)
}

func Delete(cookie string, groupName string, groupID int32, groups *models.GroupArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/usergroups.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	var err error
	if groupID != 0 {
		_, err = database.Exec("DELETE FROM user_groups WHERE user_id = $1 AND group_id = $2;", id, groupID)
	} else {
		_, err = database.Exec("DELETE FROM user_groups WHERE user_id = $1 AND group_id = (SELECT id FROM groups WHERE name = $2);", id, groupName)
	}

	if err != nil {
		log.Println("database/usergroups.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return All(cookie, groups)
}
