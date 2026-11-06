package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Request header theke Authorization token read kora
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "authorization header missing", http.StatusUnauthorized)
				return
			}

			// 1. Space diya split kora
			authParts := strings.Split(authHeader, " ")

			// 2. 2 ta kina check kora, 2 ta na hole return
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenAccess := authParts[1]

			// 3. Dot (.) diya split kora tokenAccess k
			tokenParts := strings.Split(tokenAccess, ".")

			// 4. 3 ta pailam kina check kora, na pale return
			if len(tokenParts) != 3 {
				http.Error(w, "invalid token format", http.StatusUnauthorized)
				return
			}

			headerBase64 := tokenParts[0]
			payloadBase64 := tokenParts[1]
			oldSignatureBase64 := tokenParts[2]

			// 5. Header and Payload k dot (.) diya add kora
			headerPayload := headerBase64 + "." + payloadBase64

			// 6. HeaderPayload and Secret key k byte a convert kora
			byteHeaderPayload := []byte(headerPayload)
			byteSecret := []byte(secret)

			// 7. HMAC-SHA256 diya new signature generate kora
			h := hmac.New(sha256.New, byteSecret)
			h.Write(byteHeaderPayload)
			newSignature := h.Sum(nil)

			// 8. Base64URL encode kora new signature k
			newSignatureBase64 := Base64UrlEnc(newSignature)

			// 9. Old signature and New signature check kora
			if oldSignatureBase64 != newSignatureBase64 {
				http.Error(w, "signature mismatch / token modified", http.StatusUnauthorized)
				return
			}

			// Validation successful hole porer handler a request pathano
			next.ServeHTTP(w, r)
		})
	}
}


func Base64UrlEnc(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}
