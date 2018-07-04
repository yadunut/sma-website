dev:
	go run cmd/server/main.go

.PHONY: dev

release:
	-@mkdir release 2> /dev/null || true
	packr
	GOOS=linux GOARCH=amd64 go build -o release/server ./cmd/server
	packr clean

.PHONY: release

deploy: release
	ssh yadunut@sma.yadunand.me "sudo systemctl stop sma"
	scp release/server yadunut@sma.yadunand.me:~/app/server
	ssh yadunut@sma.yadunand.me "sudo systemctl start sma"

.PHONY: release
