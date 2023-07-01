package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil{
		log.Fatal(err)
		return
	}
	w.Header().Add("Context-Type","application/json")
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, err error){
	if err != nil{
		ResponseWithJson(w,code,struct{
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}else{
		ResponseWithJson(w,code,struct{
			Error string `json:"error"`
		}{
			Error: "Internal sever error",
		})
	}
}