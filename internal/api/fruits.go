package api

import (
	"backend/internal/database/fruits"
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/valyala/fasthttp"
)

func fruitonizeProductize(f *models.Fruit, shortFruit *models.DetectedFruit) *models.ProductExtended {
	p := models.ProductExtended{}
	p.Barcode = uint64(f.ID)
	p.Name = f.NameRu
	p.Description = f.Description
	p.CategoryURL = "Fruit"
	p.Mass = strconv.FormatFloat(float64(shortFruit.Accuracy), 'f', 4, 64)
	p.BestBefore = "NULL"
	p.Manufacturer = "NULL"
	p.Ingredients = nil
	p.Image = f.Image

	var err error
	nutr := make([]*[]string, 2)
	nutr[0] = &f.NutritionLabels
	nutr[1] = &f.Nutrition
	n, err := json.Marshal(nutr)
	if err != nil {
		return nil
	}
	p.Nutrition = string(n)

	vitamins := make([]*[]string, 2)
	vitamins[0] = &f.VitaminsLabels
	vitamins[1] = &f.Vitamins
	v, err1 := json.Marshal(vitamins)
	if err1 != nil {
		return nil
	}
	p.Contents = string(v)

	return &p
}

func sendErr(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	m := &models.Msg{}
	m.Message = err.Error()
	mJSON, _ := m.MarshalJSON()
	ctx.SetBody(mJSON)
}

func Find_Fruit(ctx *fasthttp.RequestCtx) {
	log.Println("Find Fruit: " + string(ctx.Method()) + (" ") + string(ctx.Path()))
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		sendErr(ctx, err)
	}

	res := sendToPythonServer(fileHeader)
	if res == nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	shortFruit := models.DetectedFruit{}

	err = shortFruit.UnmarshalJSON(res.Body())
	if err != nil {
		sendErr(ctx, err)
	}
	fasthttp.ReleaseResponse(res)

	fruit := models.Fruit{}
	code, message := fruits.GetFruitsInfo(shortFruit.Name, &fruit)

	switch code {
	case fasthttp.StatusOK:
		prod := fruitonizeProductize(&fruit, &shortFruit)
		if prod == nil {
			sendErr(ctx, err)
		}
		json, err := prod.MarshalJSON()
		if err != nil {
			sendErr(ctx, err)
		}
		ctx.SetBody(json)
		ctx.SetStatusCode(fasthttp.StatusOK)
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
	// ctx.SetBody(res.Body())
	// ctx.SetStatusCode(fasthttp.StatusOK)
}

func sendToPythonServer(fileHeader *multipart.FileHeader) *fasthttp.Response {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in sendToPythonServer", r)
		}
	}()
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://rasseki.org:7000/predict")
	// req.SetRequestURI("http://localhost:7000/predict")
	req.Header.Add("User-Agent", "food_backend")
	// req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.SetMethod("POST")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		log.Fatal(err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)

	err = writer.Close()

	if err != nil {
		log.Fatal(err)
	}

	req.SetBody(body.Bytes())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		log.Println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		log.Println(string(bodyBytes))
	}
	fasthttp.ReleaseRequest(req)
	return resp

}
