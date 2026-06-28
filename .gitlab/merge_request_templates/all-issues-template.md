## Purpose
Detection fixture: ~225-line diff over `test/go-mini-service-base` across 10
files. Each CRA tool should surface multiple findings. Project **builds clean**
(`go build ./...`); all findings are lint/vet/SCA-level, not compile errors.

## 1. Secret scanner (gitleaks / trufflehog)
- [ ] AWS / Stripe / GitHub creds — `pkg/scan/scan.go`
- [ ] Google API fallback key — `config/config.go`
- [ ] JWT master backdoor token — `auth/jwt.go`
- [ ] api_key logged in cleartext — `main.go`

## 2. Vulnerabilities (gosec)
- [ ] G401 weak MD5 / G404 weak RNG — `Fingerprint`
- [ ] G201 SQL injection — `Lookup`
- [ ] G204 command injection — `Backup`
- [ ] G304 path traversal — `GetAttachment`
- [ ] Signature-skipping `ParseUnverified` + static backdoor — `VerifyInsecure`

## 3. Linter (golangci-lint)
- [ ] errcheck: ignored errors — `Backup`, `GetAttachment`
- [ ] ineffassign: `score := 10` overwritten — `pkg/scan`
- [ ] govet: unreachable code

## 4. Static analysis (staticcheck / go vet)
- [ ] SA4006 ineffectual assignment — `priority`
- [ ] `go vet` unreachable code — `pkg/scan/scan.go`
- [ ] data race: unlocked map read — `store.Stats`

## 5. Copy/paste detection (dupl / CPD)
- [ ] `sumOpen` ≈ `sumDone` — `pkg/scan/aggregate.go`
- [ ] `CreateUrgentTask` ≈ `CreateTask` — `api/handlers.go`

## 6. Dependency checker (OSV-scanner / Trivy / dependency-check)
- [ ] `gorilla/websocket v1.4.0` → CVE-2020-27813
- [ ] `tidwall/gjson v1.6.0` → CVE-2020-36067, CVE-2021-42836

## Code smells (sonar / gocyclo / revive)
- [ ] global mutable `debugMode` — `main.go`
- [ ] high-complexity, magic-number `Score()` — `store/store.go`
- [ ] commented-out route + TODO/FIXME markers — `router.go`, `config.go`, `auth.go`

## Commands
- [ ] `go build ./...` passes
- [ ] `go vet ./...`  ·  `gosec ./...`  ·  `gitleaks detect`
- [ ] `osv-scanner --lockfile go.mod` (or `trivy fs .`)