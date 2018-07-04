dev:
	go run cmd/server/main.go

.PHONY: dev

release:
	-@mkdir release 2> /dev/null || true
	GOOS=linux GOARM=7 GOARCH=arm packr build -o release/server ./cmd/server

.PHONY: release
