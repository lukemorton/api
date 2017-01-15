all: start

start:
	go run server/main.go

build:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/server github.com/lukemorton/api/server
	docker build -t api .

deploy: build
	now --alias api.lukemorton.co.uk

clean:
	rm -rf dist
