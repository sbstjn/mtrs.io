default: test build

clean:
	@rm -rf .build

build: clean
	@mkdir -p .build
	@go build -v -o ./.build/mtrs.io

test:
	@go test -short `go list ./... | grep -v /vendor/`

run:
	go run main.go
