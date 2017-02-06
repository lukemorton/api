all: start

test:
	go test ./...

start: test
	go run bin/server/main.go

build:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/server ./bin/server

deploy: build
	now --alias api-authors.lukemorton.co.uk

clean:
	rm -rf dist
