# go-cra-test

A small, throwaway **fixture repository** for exercising the tool flows migrated
into [`go-cra-tools`](https://gitlab.com/) (the Go re-implementation of the
Python `extra-cra` code-review agent).

The `main` branch holds clean, multi-language baseline code. Each `test/*`
branch deliberately introduces issues that **exactly one tool flow** should
flag. Open a merge request from a branch against `main`, run the agent on it, and
confirm the planted findings surface.

## Baseline languages

Go, Python, JavaScript, and Protobuf — chosen to exercise the linter's
per-language routing and the Python-only static-analysis flow.

## Branch → flow mapping

| Branch | Tool flow | What it plants | Where findings show |
|---|---|---|---|
| `test/secret-scanner` | Secret scanner (Trivy secret) | hardcoded creds in `deploy/config.env` | inline suggestions |
| `test/dependency-check` | Dependency check (Trivy vuln) | CVE'd Go modules + npm packages | overview / Security Concerns |
| `test/static-analysis` | Static analysis (ruff `--select ALL` + mypy) | Python type + lint issues | static-analysis suggestions |
| `test/cpd` | Copy-paste detection (PMD CPD) | duplicated Go logic | duplication suggestion |
| `test/linter` | Linter (gosec / ruff `-S` / eslint / buf) | per-route violations | inline suggestions |

## Notes when running the MRs

Flow enablement lives in the **agent's** config (`FeedbackConfig` / `tools:`),
not in this repo. Per the migration design, defaults are:

- Linter **ON**, secret scanner **ON**
- **CPD is OFF** — enable it for the `test/cpd` MR
- Dependency check / static analysis are **gated by language + per-tool enabled**
  — enable them for the `test/dependency-check` / `test/static-analysis` MRs

All planted secrets and vulnerable dependencies are fake / intentionally outdated
**for testing detection only**.
