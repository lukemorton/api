all: build start

build:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/server github.com/lukemorton/api/server
	docker build -t api .

start:
	docker run -it --rm --publish 3000:3000 api

deploy: build
	now --alias api.lukemorton.co.uk

clean:
	rm -rf dist
