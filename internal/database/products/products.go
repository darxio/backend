package products

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

func All(products *models.ProductArr) (code int, message string) {
	rows, err := database.Query(`SELECT barcode, name FROM products`)

	if err != nil {
		return 500, "Something went wrong.."
	}

	for rows.Next() {
		product := models.Product{}
		rows.Scan(&product.Barcode, &product.Name)
		*products = append(*products, &product)
	}
	rows.Close()

	return 200, "Successful."
}

func GetOneBarcode(barcode int64, product *models.Product) (code int, message string) {
	err := database.QueryRow(`SELECT p.barcode, p.name, array_agg(ingr.name::TEXT) AS ingredients, array_agg(ingr.type::TEXT) AS ingredient_types
	FROM products p
	JOIN product_ingredients i
	ON p.barcode = i.product_barcode
	JOIN ingredients ingr
	ON i.ingredient_id = ingr.id
	WHERE p.barcode=$1
	GROUP BY p.barcode;`, barcode).Scan(&product.Barcode, &product.Name, &product.IngredientsList, &product.IngredientTypes)


								log.Println(err)
	if err == pgx.ErrNoRows {
		return 404, "Product not found."
	} else if err != nil {
		return 500, "Something went wrong.."
		println(err.Error())
	}

	return 200, "Successful."
}
