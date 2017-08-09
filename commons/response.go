package commons

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	AppName   string      `json:"app_name"`
	Version   string      `json:"version"`
	Release   string      `json:"release"`
	Datetime  time.Time   `json:"datetime"`
	Timestamp int32       `json:"timestamp"`
	Status    string      `json:"status"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type ResponseController struct {
}

func NewResponseController() *ResponseController {
	return &ResponseController{}
}

func getResponseWithMeta(r *response, status string) *response {
	r.AppName = "tamil_font_demo"
	t := time.Now()
	r.Datetime = t
	r.Timestamp = int32(t.Unix())
	r.Status = status
	return r
}

func write(w http.ResponseWriter, response *response) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(response.Code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}

func (r ResponseController) WriteSuccess(w http.ResponseWriter, responseData interface{}, code int) {
	response := &response{
		Data:    responseData,
		Code:    code,
		Message: "ok",
	}
	response = getResponseWithMeta(response, "success")
	write(w, response)
}

func (r ResponseController) WriteError(w http.ResponseWriter, message string, code int) {
	response := &response{
		Message: message,
		Code:    code,
	}
	response = getResponseWithMeta(response, "error")
	write(w, response)
}
