# Test matrix — expected findings per flow

Each `test/*` branch (and its MR against `main`) plants issues for exactly one
go-cra-tools tool flow. Use the checklists below to verify the agent surfaces
each finding. All planted secrets and vulnerable dependencies are **fake /
intentionally outdated — for detection testing only**.

## Flow enablement (set in the agent config, not this repo)

| Flow | Default | Note |
|---|---|---|
| Secret scanner | ON | — |
| Linter | ON | — |
| Dependency check | gated | enable by language + per-tool `enabled` |
| Static analysis | gated | enable by language + per-tool `enabled` |
| CPD | **OFF** | enable `CPDEnabled` / `tools.cpd` or it won't run |

| MR | Branch | Flow | Findings appear in |
|---|---|---|---|
| !1 | `test/secret-scanner` | Secret scanner (Trivy secret) | inline suggestions |
| !2 | `test/dependency-check` | Dependency check (Trivy vuln) | overview / Security Concerns |
| !3 | `test/static-analysis` | Static analysis (ruff ALL + mypy) | static-analysis suggestions |
| !4 | `test/cpd` | Copy-paste detection (PMD CPD) | duplication suggestion |
| !5 | `test/linter` | Linter (gosec / ruff-S / eslint / buf) | inline suggestions |

---

## !1 — `test/secret-scanner`

Files: `deploy/config.env`, `deploy/id_rsa`. Findings post inline on new lines.

- [ ] AWS access key id — `AWS_ACCESS_KEY_ID` (`AKIA…`)
- [ ] AWS secret access key — `AWS_SECRET_ACCESS_KEY`
- [ ] GitHub PAT — `GITHUB_TOKEN` (`ghp_…`)
- [ ] GitLab PAT — `GITLAB_TOKEN` (`glpat-…`)
- [ ] Stripe live key — `STRIPE_SECRET_KEY` (`sk_live_…`)
- [ ] Generic password — `DB_PASSWORD`
- [ ] RSA private key block — `deploy/id_rsa`

## !2 — `test/dependency-check`

Files: `go.mod`, `go.sum`, `internal/deps/use.go`, `package.json`,
`package-lock.json`. Overview-only (Security Concerns block), not inline.

- [ ] `github.com/dgrijalva/jwt-go v3.2.0+incompatible` — CVE-2020-26160
- [ ] `gopkg.in/yaml.v2 v2.2.2` — known CVEs
- [ ] `lodash 4.17.11` — prototype pollution (CVE-2019-10744 / -2020-8203 / -2021-23337)
- [ ] `minimist 1.2.0` — prototype pollution (CVE-2020-7598 / -2021-44906)
- [ ] `axios 0.21.0` — SSRF / ReDoS (CVE-2020-28168 / -2021-3749)

## !3 — `test/static-analysis`

File: `analytics/report.py`. Importance-filtered suggestions.

_mypy (type)_
- [ ] `summarize()` returns `str` from a `-> int` function
- [ ] `to_int()` returns `str` from a `-> int` function
- [ ] `first_char()` indexes `name: str | None` (possible `None`)

_ruff (lint)_
- [ ] F401 — unused imports `os`, `sys`
- [ ] B006 — mutable default arg `numbers=[]`
- [ ] F841 — unused local `scratch`
- [ ] E711 — `total == None`
- [ ] E722 — bare `except:`
- [ ] missing docstrings / other ALL-rule noise

## !4 — `test/cpd`

Files: `pkg/orders/process_a.go`, `pkg/orders/process_b.go`. `ProcessOrderA` and
`ProcessOrderB` have token-identical bodies (~30 lines; only the name differs).

- [ ] Duplication reported between `process_a.go` and `process_b.go`

> ⚠️ CPD is OFF by default — enable it or nothing runs.

## !5 — `test/linter`

Files: `service/handler.go`, `scripts/clean.py`, `web/server.js`,
`proto/test/v1/user.proto`. Findings CWE-backfilled + second validate pass.

_Go — gosec / golangci-lint_
- [ ] G101 — hardcoded credential `apiSecret`
- [ ] G401 — weak hash `md5.Sum`
- [ ] G501 — blocklisted import `crypto/md5`
- [ ] G404 — weak RNG `math/rand`
- [ ] G204 — subprocess from variable (`exec.Command("sh","-c",userCmd)`)
- [ ] errcheck — unchecked `w.Write(...)`

_Python — ruff `S`_
- [ ] S105 — hardcoded password `PASSWORD`
- [ ] S324 — insecure `hashlib.md5`
- [ ] S602 — `subprocess.call(..., shell=True)`
- [ ] S307 — `eval(expr)`
- [ ] S101 — `assert`

_JS — eslint security_
- [ ] eval / detect-eval-with-expression — `eval(input)`
- [ ] detect-child-process — `cp.exec(userCmd)`
- [ ] detect-non-literal-fs-filename — `fs.readFileSync(name)`
- [ ] detect-object-injection — `obj[key]`

_proto — buf STANDARD_
- [ ] FIELD_LOWER_SNAKE_CASE — `userName`, `emailAddress`, `ID`
- [ ] ENUM_ZERO_VALUE_SUFFIX / ENUM_VALUE_PREFIX — `Status.UNKNOWN`, `ACTIVE`
- [ ] SERVICE_PASCAL_CASE / SERVICE_SUFFIX — `userService`
- [ ] RPC_PASCAL_CASE + RPC request/response naming — `getUser`
