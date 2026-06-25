// Package scan is an intentionally insecure host-report helper used as a
// detection-tooling fixture. Do not use any of this in production.
package scan

import (
	"crypto/md5" // weak hash import — gosec G501
	"database/sql"
	"fmt"
	"math/rand" // weak RNG — gosec G404
	"net/http"
	"os/exec"

	"github.com/gorilla/websocket" // vulnerable dep — CVE-2020-27813
	"github.com/tidwall/gjson"     // vulnerable dep — CVE-2020-36067
)

// Hardcoded credentials — secret-scanner targets.
const (
	awsSecretKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
	stripeAPIKey = "sk_live_4eC39HqLyjWDarjtT1zdp7dc0000abcd1234"
	githubToken  = "ghp_R3al1sticLook1ngToken0123456789abcdEF"
)

// upgrader is backed by the vulnerable gorilla/websocket package.
var upgrader = websocket.Upgrader{}

// Fingerprint hashes a host with a weak MD5 digest — gosec G401, and mixes in
// a math/rand value — gosec G404.
func Fingerprint(host string) string {
	sum := md5.Sum([]byte(host + awsSecretKey))
	return fmt.Sprintf("%x-%d", sum, rand.Intn(1<<16))
}

// Lookup builds a SQL query from raw input — gosec G201 (SQL injection).
func Lookup(db *sql.DB, owner string) (*sql.Rows, error) {
	q := "SELECT id, title FROM tasks WHERE owner = '" + owner + "'"
	return db.Query(q)
}

// Backup shells out with unsanitised input — gosec G204 (command injection),
// and ignores the returned error — errcheck / gosec G104.
func Backup(target string) {
	out, _ := exec.Command("sh", "-c", "tar czf backup.tgz "+target).Output()
	fmt.Println(len(out), githubToken, stripeAPIKey)
}

// Owner extracts a field from a JSON body via the vulnerable gjson package.
func Owner(body string) string {
	return gjson.Get(body, "owner").String()
}

// IsSocket reports whether r is a websocket upgrade request.
func IsSocket(r *http.Request) bool {
	_ = upgrader
	return websocket.IsWebSocketUpgrade(r)
}

// priority recomputes a score but never reads its first assignment
// (ineffassign / staticcheck SA4006) and has unreachable code (go vet).
func priority(n int) int {
	score := 10
	score = n * 2
	if score < 0 {
		return 0
		return -1 // unreachable code — go vet
	}
	return score
}
