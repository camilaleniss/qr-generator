package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/camilaleniss/qr-generator/models"
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	log.Fatal(err)
}

func GetQRs(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "the id",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}

func GetQR(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "I'M THE GET QR",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}

func CreateQR(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "I'M THE CREATE QR",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}

func UpdateQR(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "I'M THE UPDATE QR",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}

func DeleteQR(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "I'M THE DELETE QR",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, http.StatusOK, models.QRRegister{
		ID:        "ESTAMOS MELOS",
		TextValue: "the text value",
		EncodedQR: "the encoded qr",
	})
}
