package util

import (
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"
)

func GetBase64QRCode(link string) (string, error) {
	img, err := getQRCodeImage(link)
	if err != nil {
		return "", err
	}

	encodedQR, err := convertImageToBase64(img)
	if err != nil {
		return "", nil
	}

	return encodedQR, nil
}

func getQRCodeImage(textValue string) ([]byte, error) {
	var image []byte

	image, err := qrcode.Encode(textValue, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func convertImageToBase64(image []byte) (string, error) {
	imgBase64Str := base64.StdEncoding.EncodeToString(image)

	return imgBase64Str, nil
}
