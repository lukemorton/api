all: build start

build:
	docker build -t api .

start:
	docker run -it --rm --publish 3000:3000 api

deploy:
	now --alias api.lukemorton.co.uk
