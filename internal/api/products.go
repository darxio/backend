package api

import (
	"backend/internal/common"
	"backend/internal/database/products"
	"backend/internal/models"

	"log"
	"strconv"

	"github.com/valyala/fasthttp"
)

func Product_All(ctx *fasthttp.RequestCtx) {
	log.Println("Product All: " + string(ctx.Method()) + (" ") + string(ctx.Path()))

	products_ := make(models.ProductArr, 0, common.Limit)
	code, message := products.All(&products_)

	productsJSON, _ := products_.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(productsJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}

func Product_GetOneBarcode(ctx *fasthttp.RequestCtx) {
	log.Println("Product GetOneBarcode: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	barcode, _ := strconv.ParseInt(ctx.UserValue("barcode").(string), 10, 64)

	p := &models.Product{}
	code, message := products.GetOneBarcode(int64(barcode), p)

	pJSON, _ := p.MarshalJSON()

	switch code {
	case fasthttp.StatusOK:
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetBody(pJSON)
	case fasthttp.StatusNotFound:
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	case fasthttp.StatusInternalServerError:
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		m := &models.Msg{}
		m.Message = message
		mJSON, _ := m.MarshalJSON()
		ctx.SetBody(mJSON)
	}
}
