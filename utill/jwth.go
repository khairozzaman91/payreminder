package utill

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub       int    `json:"sub"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func CreateJwt(secret string, data Payload) (string, error) {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	byteHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	bytePayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	headerBase64 := Base64UrlEnc(byteHeader)
	payloadBase64 := Base64UrlEnc(bytePayload)

	headerPayload := headerBase64 + "." + payloadBase64

	byteHeaderPayload := []byte(headerPayload)
	byteSecret := []byte(secret)

	h := hmac.New(sha256.New, byteSecret)
	h.Write(byteHeaderPayload)

	signature := h.Sum(nil)

	signatureBase64 := Base64UrlEnc(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64

	return jwt, nil
}

func Base64UrlEnc(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}