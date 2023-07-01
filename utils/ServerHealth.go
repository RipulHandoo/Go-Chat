package utils

import (
	"fmt"
	"net/http"

	"github.com/RipulHandoo/goChat/db"
)

type resp struct{
	Status string
}

func ServerHealth(w http.ResponseWriter, r *http.Request) {
	db := db.DBInstance()

	if db == nil{
		ResponseWithError(w, http.StatusBadRequest,fmt.Errorf("database error"))
		return
	}
	ResponseWithJson(w,200,resp{
		Status: "OK",
	})
}