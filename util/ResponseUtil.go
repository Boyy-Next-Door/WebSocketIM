package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

func Ok(w http.ResponseWriter, data interface{}, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	marshal, _ := json.Marshal(Response{msg, data, 200})
	w.Write(marshal)
}

func ParamError(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	marshal, _ := json.Marshal(Response{"parameter error : " + err, nil, 400})
	w.Write(marshal)
}

func InternalError(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	marshal, _ := json.Marshal(Response{"internal error : " + err, nil, 500})
	w.Write(marshal)
}
