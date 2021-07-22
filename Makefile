clean:
	. ./scripts/clean.sh
build:
	. ./scripts/build.sh
start:
	. ./scripts/start.sh
dev:
	go run -x ./cmd/email-service/main.go
format:
	go fmt ./...
mockgen:
	go generate ./...
test:
	go test ./...
testcov:
	go test ./... -cover -v -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
.PHONY:
	build
