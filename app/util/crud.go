package util

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/camilaleniss/qr-generator/app/models"
)

const (
	QRTABLE = "qrcodes"

	QUERIESLIMIT = 20

	SELECTALLQUERY = "SELECT * FROM %s LIMIT %d;"

	INSERTQUERY = "INSERT INTO %s(text_value, encoded_qr) VALUES(?, ?);"

	SELECTBYIDQUERY = "SELECT * FROM %s WHERE id=?;"

	DELETEBYIDQUERY = "DELETE FROM %s WHERE id=?;"

	UPDATEBYIDQUERY = "UPDATE %s SET text_value=?, encoded_qr=? WHERE id=?;"
)

var (
	ErrNotFoundRegister = errors.New("not found register in table")
)

// SelectALLQRs gets all the QR Registers in table
func SelectAllQRs(db *sql.DB) ([]models.QRRegister, error) {
	query := fmt.Sprintf(SELECTALLQUERY, QRTABLE, QUERIESLIMIT)

	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	qrCodes := []models.QRRegister{}

	for result.Next() {
		qrRegister := models.QRRegister{}

		err := result.Scan(&qrRegister.ID, &qrRegister.TextValue, &qrRegister.EncodedQR)
		if err != nil {
			return nil, err
		}

		qrCodes = append(qrCodes, qrRegister)
	}

	return qrCodes, nil
}

// SelectQRByID gets a QR Register with a given ID
func SelectQRByID(db *sql.DB, id int) (models.QRRegister, error) {
	query := fmt.Sprintf(SELECTBYIDQUERY, QRTABLE)

	result, err := db.Query(query, id)
	if err != nil {
		return models.QRRegister{}, err
	}

	defer result.Close()

	qrRegister := models.QRRegister{}

	for result.Next() {
		err := result.Scan(&qrRegister.ID, &qrRegister.TextValue, &qrRegister.EncodedQR)
		if err != nil {
			return models.QRRegister{}, err
		}

		return qrRegister, nil
	}

	return models.QRRegister{}, ErrNotFoundRegister
}

// CreateQRRegister inserts a new QR register in the qrcodes table
func CreateQRRegister(db *sql.DB, textValue, encodedQR string) (int, error) {
	queryString := fmt.Sprintf(INSERTQUERY, QRTABLE)

	query, err := db.Prepare(queryString)
	if err != nil {
		return 0, err
	}

	response, err := query.Exec(textValue, encodedQR)
	if err != nil {
		return 0, err
	}

	id, err := response.LastInsertId()

	return int(id), err
}

// UpdateQRRegister updates a QR Register values with a given id
func UpdateQRRegister(db *sql.DB, id int, textValue, encodedQR string) error {
	queryString := fmt.Sprintf(UPDATEBYIDQUERY, QRTABLE)

	query, err := db.Prepare(queryString)
	if err != nil {
		return err
	}

	result, err := query.Exec(textValue, encodedQR, id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if int(count) == 0 {
		return ErrNotFoundRegister
	}

	return nil
}

// DeleteQRRegister deletes a QR Register with a given id
func DeleteQRRegister(db *sql.DB, id int) error {
	queryString := fmt.Sprintf(DELETEBYIDQUERY, QRTABLE)

	query, err := db.Prepare(queryString)
	if err != nil {
		return err
	}

	result, err := query.Exec(id)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if int(count) == 0 {
		return ErrNotFoundRegister
	}

	return nil
}
