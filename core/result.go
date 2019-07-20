package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	RetCode int32
	Data interface{}
	Message string
}

func ResultOk(w http.ResponseWriter, data interface{})  {
	response := Response{
		RetCode: 0,
		Data: data,
		Message: "success",
	}
	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}

func ResultFail(w http.ResponseWriter, err string)  {
	response := Response{
		RetCode: -1,
		Data: err,
		Message: "success",
	}
	responseBytes, _ := json.Marshal(response)
	w.Write(responseBytes)
}
