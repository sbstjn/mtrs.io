default: build

clean:
	rm -rf .build

build: clean
	mkdir -p .build
	go build -v -o ./.build/mtrs.io

run:
	go run main.go
