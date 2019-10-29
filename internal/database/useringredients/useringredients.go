package useringredients

/*

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

func AllExcludedIngredients(cookie string, ingredients *models.IngredientArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/useringredients.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	rows, err := database.Query(`SELECT id, name, about
								FROM ingredients
								JOIN excluded_ingredients
								ON ingredients.id = excluded_ingredients.ingredient_id
								WHERE user_id = $1;`, id)

	if err != nil {
		log.Println("database/useringredients.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		ingredient := models.Ingredient{}
		rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Type)
		*ingredients = append(*ingredients, &ingredient)
	}

	rows.Close()

	return 200, "Successful."
}

func AddExcludedIngredient(cookie string, ingredientName string, ingredientID int32, ingredients *models.IngredientArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/useringredients.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	var err error
	if ingredientID != 0 {
		_, err = database.Exec("INSERT INTO excluded_ingredients(user_id, ingredient_id) VALUES ($1, $2);", id, ingredientID)
	} else {
		_, err = database.Exec("INSERT INTO excluded_ingredients(user_id, ingredient_id) VALUES ($1, (SELECT id FROM ingredients WHERE name = $2));", id, ingredientName)
	}

	if err != nil {
		pgErr := err.(pgx.PgError)
		if pgErr.Code == "23505" {
			log.Println("database/useringredients.go: 409, " + err.Error())
			return 409, "This user has already excluded this ingredient."
		}
		if pgErr.Code == "23503" {
			log.Println("database/useringredients.go: 404, " + err.Error())
			return 404, "This ingredient doesn't exist."
		}
		log.Println("database/useringredients.go: 500, " + err.Error())
		return 500, "Something went wrong.."
	}

	return AllExcludedIngredients(cookie, ingredients)
}

func DeleteExcludedIngredient(cookie string, ingredientName string, ingredientID int32, ingredients *models.IngredientArr) (code int, message string) {
	var id int
	errS := database.QueryRow(`SELECT user_id FROM sessions WHERE cookie = $1`, cookie).Scan(&id)

	if errS != nil {
		log.Println("database/useringredients.go: 500, " + errS.Error())
		return 500, errS.Error()
	}

	var err error
	if ingredientID != 0 {
		_, err = database.Exec("DELETE FROM excluded_ingredients WHERE user_id = $1 AND ingredient_id = $2;", id, ingredientID)
	} else {
		_, err = database.Exec("DELETE FROM excluded_ingredients WHERE user_id = $1 AND ingredient_id = (SELECT id FROM ingredients WHERE name = $2);", id, ingredientName)
	}

	if err != nil {
		log.Println("database/useringredients.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return AllExcludedIngredients(cookie, ingredients)
}
*/
