package main

import (
	"net/http"

	"github.com/camilaleniss/qr-generator/controller"
	"github.com/camilaleniss/qr-generator/util"
	"github.com/gorilla/mux"
)

func main() {
	db := util.GetConnection()

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/qrcodes", controller.GetQRs).Methods("GET")
	router.HandleFunc("/qrcodes", controller.CreateQR).Methods("POST")
	router.HandleFunc("/qrcodes/health", controller.GetHealth).Methods("GET")
	router.HandleFunc("/qrcodes/{id}", controller.GetQR).Methods("GET")
	router.HandleFunc("/qrcodes/{id}", controller.UpdateQR).Methods("PUT")
	router.HandleFunc("/qrcodes/{id}", controller.DeleteQR).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
