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
.PHONY:
	build