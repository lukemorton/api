all: build run

build:
	docker build -t api .

run:
	docker run -it --rm --publish 3000:3000 api

deploy:
	now --alias api.lukemorton.co.uk
