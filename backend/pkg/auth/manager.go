package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type TokenManager interface {
	NewJWT(userId string, ttl time.Duration) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
	MiddlewareJWT(next http.Handler) http.Handler
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userId string, userRole, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(ttl).Unix(),
		"sub":  userId,
		"role": userRole,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (string, int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, fmt.Errorf("error get user claims from token")
	}

	userId := claims["sub"].(string)
	userRole := int(claims["role"].(float64)) // JWT по умолчанию декодирует числа как float64

	return userId, userRole, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (m *Manager) MiddlewareJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("Authorization")
		if tokenString != "" {
			tokenString = tokenString[len("Bearer "):]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(m.signingKey), nil
			})
			if err != nil || !token.Valid {
				http.Error(writer, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(writer, "Unauthorized", http.StatusUnauthorized)
				return
			}

			userId, userRole := claims["sub"].(string), int(claims["role"].(float64))

			// Установите user ID и role в контекст запроса
			ctx := context.WithValue(request.Context(), "userID", userId)
			ctx = context.WithValue(ctx, "userRole", userRole)
			request = request.WithContext(ctx)

			next.ServeHTTP(writer, request)
		} else {
			http.Error(writer, "Unauthorized: No token in the header", http.StatusUnauthorized)
		}
	})
}
