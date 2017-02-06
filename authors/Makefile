all: start

test:
	go test ./...

start: test
	go run bin/server/main.go

build: test
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/server ./bin/server
	docker build -t api-authors .

docker: build
	docker run -it --rm --publish 3000:3000 api-authors

deploy: build
	now --alias api-authors.lukemorton.co.uk

clean:
	rm -rf dist
