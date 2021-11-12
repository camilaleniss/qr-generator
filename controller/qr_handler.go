package controller

import (
	"encoding/json"
	"net/http"

	"github.com/camilaleniss/qr-generator/models"
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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
