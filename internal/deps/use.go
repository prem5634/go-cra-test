// Package deps references intentionally outdated, vulnerable dependencies so the
// dependency-check tool flow has something to flag. TEST FIXTURE ONLY.
package deps

import (
	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

// LoadConfig unmarshals YAML config using an outdated gopkg.in/yaml.v2.
func LoadConfig(data []byte) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ParseToken parses (without verifying) a JWT using an outdated jwt-go.
func ParseToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	parser := new(jwt.Parser)
	if _, _, err := parser.ParseUnverified(token, claims); err != nil {
		return nil, err
	}
	return claims, nil
}
