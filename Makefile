clean:
	. ./scripts/clean.sh
build:
	. ./scripts/build.sh
start:
	. ./scripts/start.sh
start-dev:
	go run ./cmd/email-service/main.go
.PHONY:
	build