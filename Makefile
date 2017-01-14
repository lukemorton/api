all: build run

build:
	docker build -t api .

run:
	docker run -it --rm api

deploy:
	now --alias api.lukemorton.co.uk
