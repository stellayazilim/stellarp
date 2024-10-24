SHELL := powershell.exe

export GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/postgres
export GOOSE_DRIVER=postgres
export GOOSE_MIGRATION_DIR=./migrations


run\:dev:
	wgo run ./cmd/

run\:test:
	go test ./... ./modules/config/...  ./modules/identity/...

run\:test\:coverage:
	go test ./... ./modules/config/... ./modules/identity/...   -coverprofile cover.out -coverpkg ./... ./modules/config/... ./modules/identity/...
run\:test\:watch:
	go run github.com/onsi/ginkgo/v2/ginkgo watch  -p -r -v  ./modules/config/...  ./modules/identity/...
migrate\:up:
	goose up;

migrate\:down:
	goose down;

migrate\:create:
	goose create $(name) sql

