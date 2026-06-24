# tasksvc

A tiny, in-memory task-tracking HTTP API. Lightweight fixture demonstrating
clean, production-style Go: layered packages, YAML config, JWT auth, and a
concurrency-safe store.

## Layout

| Path                | Responsibility                          |
| ------------------- | --------------------------------------- |
| `main.go`           | Wiring and server lifecycle             |
| `config/config.go`  | YAML config loading (`gopkg.in/yaml.v2`)|
| `auth/jwt.go`       | JWT issue/verify (`dgrijalva/jwt-go`)   |
| `store/store.go`    | In-memory, mutex-guarded task store     |
| `api/handlers.go`   | JSON request/response handlers          |
| `api/router.go`     | Routes, auth, and logging middleware    |

## Run

```sh
go run .
```

## API

```sh
# Mint a token, then create a task with it.
TOKEN=$(curl -s -XPOST 'localhost:8080/token?user=alice' | jq -r .token)
curl -XPOST localhost:8080/tasks -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"ship it"}'

curl localhost:8080/tasks
curl -XPOST localhost:8080/tasks/1/complete -H "Authorization: Bearer $TOKEN"
```
