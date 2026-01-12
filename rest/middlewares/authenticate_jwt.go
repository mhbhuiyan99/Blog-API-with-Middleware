package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

// Adding struct for getting login user_id to create post
type Payload struct {
	Sub      int    `json:"sub"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsWriter bool   `json:"is_writer"`
}

/*
1. parse jwt
2. parse header and payload or claims
3. hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
4. parse signature part from the jwt
5. if the signature and hash is same -> forword to create post
6. otherwise 401 status code with Unauthorized
*/

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// get token by splitting bearer and token
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]

		// split the JWT token into three parts
		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JwtSecretKey)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)

		if signature != newSignature {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Decode payload to get user info for further use like create post,...
		payloadBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(jwtPayload)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var payload Payload
		err = json.Unmarshal(payloadBytes, &payload)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Store user_id in context for further use
		ctx := context.WithValue(r.Context(), "user_id", payload.Sub)


		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
