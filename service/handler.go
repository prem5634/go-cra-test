// Package service is an intentionally insecure HTTP handler for linter testing
// (gosec / golangci-lint). Do not use any of this in real code.
package service

import (
	"crypto/md5" // weak hash import — gosec G501
	"fmt"
	"math/rand" // weak RNG — gosec G404
	"net/http"
	"os/exec"
)

// apiSecret is a hardcoded credential — gosec G101.
const apiSecret = "AKIA2E0A8F3B5C7D9G1Hsecretkeymaterial"

// Hash returns an MD5 hex digest of data — gosec G401 (weak cryptographic hash).
func Hash(data string) string {
	sum := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", sum)
}

// Token returns a "random" token using math/rand — gosec G404 (weak RNG).
func Token() int {
	return rand.Intn(1000000)
}

// Run executes a shell command built from user input — gosec G204 (subprocess
// launched with a variable / command injection).
func Run(userCmd string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", userCmd)
	return cmd.Output()
}

// Handler writes a response but ignores the write error — errcheck.
func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("secret=" + apiSecret))
}
