package api

import (
	"backend/internal/models"
	"bytes"
	"io"
	"log"
	"mime/multipart"

	"github.com/valyala/fasthttp"
)

func sendErr(ctx *fasthttp.RequestCtx, err error) {
	ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	m := &models.Msg{}
	m.Message = err.Error()
	mJSON, _ := m.MarshalJSON()
	ctx.SetBody(mJSON)
}

func Find_Fruit(ctx *fasthttp.RequestCtx) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		sendErr(ctx, err)
	}

	sendToPythonServer(fileHeader)

	ctx.SetStatusCode(fasthttp.StatusOK)
	// ctx.SetBody(string(res.Body()))

}

func sendToPythonServer(fileHeader *multipart.FileHeader) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://localhost:5000/predict")
	req.Header.Add("User-Agent", "food_backend")
	// req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.SetMethod("POST")

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// defer recover() {}
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
	fasthttp.ReleaseResponse(resp)
	return resp

}
