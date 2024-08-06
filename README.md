# COMANDOS

# AULA 01
- go mod init [nome-repositório]
- go install github.com/jackc/tern/v2@latest
- tern init ./internal/store/pgstore/migrations
- tern new --migrations ./internal/store/pgstore/migrations [nome-migration]
- go mod tidy
- go run cmd/tools/terndotenv/main.go
- go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
- sqlc generate -f [path]
- go generate ./...
