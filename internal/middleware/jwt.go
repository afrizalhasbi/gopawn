package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// type Claims struct {
// 	Uuid string
// 	jwt.RegisteredClaims
// }

type JwtMiddleware struct {
	SecretKey []byte
}

func (m *JwtMiddleware) ValidateJwt(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := r.Header.Get("X-User-ID")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || uuid == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		jwtTokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// jwt.ParseWithClaims(tokenString string, claims jwt.Claims, keyFunc jwt.Keyfunc, options ...jwt.ParserOption)

		jwtToken, err := jwt.Parse(jwtTokenString, m.VerifyParsedToken)

		if err != nil {
			http.Error(w, "Unauthorized", 401)
			return
		}

		if !jwtToken.Valid {
			http.Error(w, "Unauthorized", 401)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (m *JwtMiddleware) VerifyParsedToken(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Unauthorized")
	} else {
		return m.SecretKey, nil
	}
}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString(m.SecretKey)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to sign token: %w", err)
// 	}

// 	return tokenString, nil
// }
