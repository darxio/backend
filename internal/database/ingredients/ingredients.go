package ingredients

import (
	"backend/internal/database/connection"
	"backend/internal/models"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

func All(ingredients *models.IngredientArr) (code int, message string) {
	rows, err := database.Query("SELECT id, name, about FROM ingredients")

	if err != nil {
		return 500, "Something went wrong.."
	}

	for rows.Next() {
		ingredient := models.Ingredient{}
		rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.About)
		*ingredients = append(*ingredients, &ingredient)
	}
	rows.Close()

	return 200, "Successful."
}

func About(ingredientName string, ingredientID int32, ingredient *models.Ingredient) (code int, message string) {
	var err error
	if ingredientID != 0 {
		err = database.QueryRow("SELECT id, name, about FROM ingredients WHERE id = $1;", ingredientID).Scan(&ingredient.ID, &ingredient.Name, &ingredient.About)
	} else {
		err = database.QueryRow("SELECT id, name, about FROM ingredients WHERE name = $1;", ingredientName).Scan(&ingredient.ID, &ingredient.Name, &ingredient.About)
	}

	if err == pgx.ErrNoRows {
		return 404, "Ingredient not found."
	} else if err != nil {
		return 500, "Something went wrong.."
	}

	return 200, "Successful."
}

func GroupAll(ingredientName string, ingredientID int32, ingredients *models.IngredientArr) (code int, message string) {

	// if err == pgx.ErrNoRows {
	// 	return 404, "Ingredient not found."
	// } else if err != nil {
	// 	return 500, "Something went wrong.."
	// }

	return 200, "Successful"
}
