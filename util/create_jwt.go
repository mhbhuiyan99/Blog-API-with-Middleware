package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub            int    `json:"sub"` // user id
	Username       string `json:"username"`
	Email          string `json:"email"`
	IsWriter       bool   `json:"is_writer"`
	ExpirationTime int64  `json:"expiration_time"`
	IssuedAt       int64  `json:"issued_at"`
}

func CreateJWT(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	// Set expiration time (7 days)
	data.ExpirationTime = time.Now().Add(7 * 24 * time.Hour).Unix()
	data.IssuedAt = time.Now().Unix()

	// Marshal to convert into byte array
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	// convert into base64 string
	headerB64 := base64UrlEncode(byteArrHeader)

	byteArrPayload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(byteArrPayload)

	message := headerB64 + "." + payloadB64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
