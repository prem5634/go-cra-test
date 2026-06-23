// Package service is an intentionally insecure auth helper for linter testing
// (gosec / golangci-lint). Do not use any of this in real code.
package service

import (
	"crypto/md5" // weak hash import — gosec G501
	"fmt"
	"math/rand" // weak RNG — gosec G404
	"net/http"
	"os/exec"
)

// adminToken is a hardcoded credential — gosec G101.
const adminToken = "glpat-Zz9Yy8Xx7Ww6Vv5Uu4T"

// Fingerprint returns an MD5 hex digest — gosec G401 (weak cryptographic hash).
func Fingerprint(data string) string {
	sum := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", sum)
}

// SessionID returns a "random" session id using math/rand — gosec G404 (weak RNG).
func SessionID() int {
	return rand.Intn(1000000)
}

// Exec runs a shell command built from user input — gosec G204 (command injection).
func Exec(userCmd string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", userCmd)
	return cmd.Output()
}

// Login writes a response but ignores the write error — errcheck.
func Login(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("token=" + adminToken))
}
