# api.lukemorton.co.uk/users

## Prerequisites

- Install [go](https://golang.org/dl/) :D **required**
- Install [now-cli](https://github.com/zeit/now-cli/) if you're deploying to now
- Install [docker](https://www.docker.com/products/overview#/install_the_platform) if you're testing the docker file locally

## Developing

**Installing:**

If you have go locally install you can just use `go get`:

```
cd $GOPATH
go get github.com/lukemorton/api/users
cd src/github.com/lukemorton/api/users
go get ./...
```

**Setup database:**

Create MySQL instance with docker:

```
docker run --name api-users-mysql -e MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=users -d mysql
export DATABASE_URL="root:root@/users?parseTime=true&clientFoundRows=true"
```

**Running tests:**

```
make test
```

**Running app:**

```
make
```

This will boot the app on [http://localhost:3000](http://localhost:3000)

**Running app inside docker:**

As this is what happens in production when the app is deployed it's worth testing.

```
make docker
```
