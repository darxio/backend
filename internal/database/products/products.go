package products

import (
	"backend/internal/database/connection"
	"backend/internal/models"
	"log"
	"strconv"

	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

/*
func All(products *models.ProductExtendedArr) (code int, message string) {
	rows, err := database.Query(`
		SELECT
		barcode,
		name,
		description,
		contents,
		category_url,
		mass,
		bestbefore,
		nutrition,
		manufacturer,
		image
		FROM products_extended`)

	if err != nil {
		log.Println("database/products.go: 500, " + err.Error())
		return 500, err.Error()
	}

	for rows.Next() {
		product := models.ProductExtended{}
		rows.Scan(&product.Barcode, &product.Name, &product.Description, &product.Contents, &product.CategoryURL, &product.Mass, &product.BestBefore, &product.Nutrition, &product.Manufacturer, &product.Image)
		*products = append(*products, &product)
	}
	rows.Close()

	return 200, "Successful."
}
*/
func GetOneBarcode(barcode int64, productExt *models.ProductExtended,
	productShr *models.ProductShrinked, shrinked *bool) (code int, message string) {
	err := database.QueryRow(`
	SELECT
		barcode,
		COALESCE(name, 'NULL'),
		COALESCE(description, 'NULL'),
		COALESCE(contents, 'NULL'),
		COALESCE(category_url, 'NULL'),
		COALESCE(mass, 'NULL'),
		COALESCE(bestbefore, 'NULL'),
		COALESCE(nutrition, 'NULL'),
		COALESCE(manufacturer, 'NULL'),
		COALESCE(image, 'NULL')
		FROM products_extended
		WHERE barcode = $1;`, barcode).Scan(
		&productExt.Barcode, &productExt.Name,
		&productExt.Description, &productExt.Contents,
		&productExt.CategoryURL, &productExt.Mass,
		&productExt.BestBefore, &productExt.Nutrition,
		&productExt.Manufacturer, &productExt.Image)

	productExt.Image = "http://www.goodsmatrix.ru/BigImages/" + strconv.FormatUint(productExt.Barcode, 10) + ".jpg"

	if err == pgx.ErrNoRows {
		errSelect := database.QueryRow(`SELECT barcode, name FROM products WHERE barcode = $1;`, barcode).Scan(&productShr.Barcode, &productShr.Name)

		if errSelect != nil {
			if errSelect == pgx.ErrNoRows {
				log.Println("database/products.go: 404, " + err.Error())
				return 404, "Product not found."
			}
			log.Println("database/products.go: 500, " + err.Error())
			return 500, err.Error()
		}

		*shrinked = true
		return 200, "Successful."

	} else if err != nil {
		log.Println("database/products.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "Successful."
}

func GetManyByName(name string, productExt *models.ProductExtendedArr,
	productShr *models.ProductShrinkedArr, shrinked *bool) (code int, message string) {
	res, err := database.Query(`
	SELECT
		barcode,
		COALESCE(name, 'NULL'),
		COALESCE(description, 'NULL'),
		COALESCE(contents, 'NULL'),
		COALESCE(category_url, 'NULL'),
		COALESCE(mass, 'NULL'),
		COALESCE(bestbefore, 'NULL'),
		COALESCE(nutrition, 'NULL'),
		COALESCE(manufacturer, 'NULL'),
		COALESCE(image, 'NULL')
		FROM products_extended
		WHERE name LIKE '%'||$1||'%' AND category_url LIKE '%Товары/Продукты питания%'
		ORDER BY length(description)
		LIMIT 10`, name)
	if err == nil && res.Err() == nil {
		for res.Next() {
			curProd := models.ProductExtended{}
			res.Scan(
				&curProd.Barcode, &curProd.Name,
				&curProd.Description, &curProd.Contents,
				&curProd.CategoryURL, &curProd.Mass,
				&curProd.BestBefore, &curProd.Nutrition,
				&curProd.Manufacturer, &curProd.Image)
			curProd.Image = "http://www.goodsmatrix.ru/BigImages/" + strconv.FormatUint(curProd.Barcode, 10) + ".jpg"
			*productExt = append(*productExt, &curProd)
		}
	} else if err == pgx.ErrNoRows {
		rows, errSelect := database.Query(`SELECT barcode, name FROM products WHERE name like '%'|| $1 ||'%';`, name)

		if errSelect == nil && rows.Err() == nil {
			for rows.Next() {
				curProd := models.ProductShrinked{}
				rows.Scan(&curProd.Barcode, &curProd.Name)
				*productShr = append(*productShr, &curProd)
			}
		}

		if errSelect != nil {
			if errSelect == pgx.ErrNoRows {
				log.Println("database/products.go: 404, " + err.Error())
				return 404, "Product not found."
			}
			log.Println("database/products.go (shrinked): 500, " + err.Error())
			return 500, err.Error()
		}

		*shrinked = true

		return 200, "Successful."

	} else if err != nil {
		log.Println("database/products.go: 500, " + err.Error())
		return 500, err.Error()
	}

	return 200, "Successful."
}

func Add(barcode int64, name string) (code int, message string) {
	_, err := database.Exec("INSERT INTO moderation_products(barcode, name) VALUES ($1, $2);", barcode, name)
	if err != nil {
		log.Println("database/products.go (shrinked): 500, " + err.Error())
		return 500, err.Error()
	}
	return 200, "Successful."

}
