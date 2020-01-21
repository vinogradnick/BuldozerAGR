package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitServer() {
	router := mux.NewRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Print(err)
	}
}
