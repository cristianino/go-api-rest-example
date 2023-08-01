package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status        int         `json:"status"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	contentType   string
	responseWrite http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:        http.StatusOK,
		responseWrite: rw,
		contentType:   "application/json",
	}
}

func (resp *Response) Send() {
	resp.responseWrite.Header().Set("Content-Type", resp.contentType)
	resp.responseWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.responseWrite, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (res *Response) NotFound() {
	res.Status = http.StatusNotFound
	res.Message = "Resource not found."
}

func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NotFound()
	response.Send()
}

func (res *Response) UnprocessableEntity() {
	res.Status = http.StatusUnprocessableEntity
	res.Message = "Unprocessable Request."
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
