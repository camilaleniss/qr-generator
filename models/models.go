package models

type QRRegister struct {
	ID        string `json:"id"`
	TextValue string `json:"text_value"`
	EncodedQR string `json:"encoded_qr"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
