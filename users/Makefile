ENV := DATABASE_URL="root:root@/users?parseTime=true&clientFoundRows=true"

all: start

db:
ifeq ($(shell docker ps -a --format "{{.Names}}" | grep api-users-mysql-for-host),)
	docker run -d --name api-users-mysql-for-host --publish 3306:3306 --env MYSQL_ROOT_PASSWORD=root --env MYSQL_DATABASE=users mysql
	# Sleep for a minute to wait for MySQL to setup
	sleep 60
else
	@echo "MySQL container api-users-mysql-for-host exists"
endif

test: db
	$(ENV) go test .
	$(ENV) go test ./bin/server

start: test
	$(ENV) go run bin/server/main.go

build:
	mkdir -p dist
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/server ./bin/server
	docker build -t api-users .

docker: build
	docker run -it --rm --publish 3000:3000 api

deploy: build
	now --alias api-users.lukemorton.co.uk

clean:
	rm -rf dist
