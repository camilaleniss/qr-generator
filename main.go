package main

import (
	"database/sql"
	"net/http"

	"github.com/camilaleniss/qr-generator/controller"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/QR_CODE_SYSTEM")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/qrcodes", controller.GetQRs).Methods("GET")
	router.HandleFunc("/qrcodes", controller.CreateQR).Methods("POST")
	router.HandleFunc("/qrcodes/{id}", controller.GetQR).Methods("GET")
	router.HandleFunc("/qrcodes/{id}", controller.UpdateQR).Methods("PUT")
	router.HandleFunc("/qrcodes/{id}", controller.DeleteQR).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
