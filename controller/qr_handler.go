package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/camilaleniss/qr-generator/models"
	"github.com/camilaleniss/qr-generator/util"
	"github.com/gorilla/mux"
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
	if err != nil {
		log.Fatal(err)
	}
}

func GetQRs(w http.ResponseWriter, r *http.Request) {
	db := util.GetConnection()

	defer db.Close()

	qrCodes, err := util.SelectAllQRs(db)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	respondwithJSON(w, http.StatusOK, qrCodes)
}

func GetQR(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	db := util.GetConnection()

	defer db.Close()

	qrCode, err := util.SelectQRByID(db, id)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	respondwithJSON(w, http.StatusOK, qrCode)
}

func CreateQR(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	keyVal := make(map[string]string)

	err = json.Unmarshal(body, &keyVal)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	textValue := keyVal["text_value"]

	encodedQR, err := util.GetBase64QRCode(textValue)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	db := util.GetConnection()

	defer db.Close()

	id, err := util.CreateQRRegister(db, textValue, encodedQR)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})

		return
	}

	respondwithJSON(w, http.StatusOK, map[string]string{
		"Status": "Created",
		"QR ID":  fmt.Sprintf("%d", id),
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
