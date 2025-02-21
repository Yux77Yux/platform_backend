package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT Secret
var jwtSecret = []byte("your_secret_key")

// Claims represents the JWT claims
type Claims struct {
	UserID int64    `json:"user_id"`
	Role   string   `json:"role"`
	Scope  []string `json:"scope"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a user with a given role and scope
func GenerateJWT(userID int64, role string, scope []string) (string, error) {
	// Set expiration time based on whether it's an access token or refresh token
	expirationTime := time.Now().Add(30 * time.Minute) // default for access token
	NoScope := len(scope) == 0
	if NoScope {
		// No scope means refresh token
		expirationTime = time.Now().Add(7 * 24 * time.Hour) // Refresh token expires in 7 days
	}

	claims := &Claims{
		UserID: userID,
		Role:   role,
		Scope:  scope,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	if NoScope {
		str, err = EncryptRefreshToken(str)
		if err != nil {
			return "", err
		}
	}
	return str, nil
}

func ParseJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		// 细化错误信息
		if strings.Contains(err.Error(), "unexpected signing method") {
			return nil, errors.New("INVALID_SIGNATURE_METHOD")
		} else if strings.Contains(err.Error(), "malformed") || strings.Contains(err.Error(), "invalid number of segments") {
			return nil, errors.New("INVALID_TOKEN_FORMAT")
		}
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("INVALID_TOKEN")
	}

	return claims, nil
}
