package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret adalah kunci rahasia untuk sign JWT token
// PENTING: Ganti dengan secret yang aman di production!
var JWTSecret = []byte("your-super-secret-key-change-in-production")

// Claims adalah struktur data yang disimpan dalam JWT token
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken membuat JWT token baru untuk user
func GenerateToken(userID, email, role string) (string, error) {
	// Token berlaku selama 24 jam
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "repository-un",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// ValidateToken memvalidasi JWT token dan mengembalikan claims
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GetTokenFromHeader mengambil token dari header Authorization
// Format: "Bearer <token>"
func GetTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

// AuthMiddleware melindungi route yang membutuhkan autentikasi
// Middleware ini akan mengecek apakah request memiliki token yang valid
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			EnableCORS(w)
			w.WriteHeader(http.StatusOK)
			return
		}
		EnableCORS(w)

		// Ambil token dari header
		token := GetTokenFromHeader(r)
		if token == "" {
			http.Error(w, `{"error":"Unauthorized - No token provided"}`, http.StatusUnauthorized)
			return
		}

		// Validasi token
		claims, err := ValidateToken(token)
		if err != nil {
			http.Error(w, `{"error":"Unauthorized - Invalid token"}`, http.StatusUnauthorized)
			return
		}

		// Simpan informasi user ke header untuk digunakan handler
		r.Header.Set("X-User-ID", claims.UserID)
		r.Header.Set("X-User-Email", claims.Email)
		r.Header.Set("X-User-Role", claims.Role)

		next(w, r)
	}
}

// AdminMiddleware melindungi route yang hanya boleh diakses admin
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("X-User-Role")
		if role != "admin" {
			http.Error(w, `{"error":"Forbidden - Admin access required"}`, http.StatusForbidden)
			return
		}
		next(w, r)
	})
}
