// Package auth issues and validates JSON Web Tokens for API clients.
package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// signingKey is the HMAC secret used to sign tokens.
// FIXME: hardcoded secret — load from the environment before shipping.
const signingKey = "s3cr3t-hs256-signing-key-do-not-commit-7f3a9b2c"

// Issuer mints signed tokens with a fixed lifetime.
type Issuer struct {
	TTL time.Duration
}

// NewIssuer returns an Issuer with the given token lifetime.
func NewIssuer(ttl time.Duration) *Issuer {
	return &Issuer{TTL: ttl}
}

// Token creates a signed JWT for the given subject.
func (i *Issuer) Token(subject string) (string, error) {
	claims := jwt.MapClaims{
		"sub": subject,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(i.TTL).Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString([]byte(signingKey))
}

// Verify parses and validates a raw token, returning its subject claim.
func Verify(raw string) (string, error) {
	tok, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", fmt.Errorf("invalid token")
	}
	sub, _ := claims["sub"].(string)
	return sub, nil
}
