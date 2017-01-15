# api.lukemorton.co.uk

## Prerequisites

- Install [docker](https://www.docker.com/products/overview#/install_the_platform) :)
- Install [go](https://golang.org/dl/) :D
- Install [now-cli](https://github.com/zeit/now-cli/) if you're deploying to now

## Developing

**Installing:**

If you have go locally install you can just use `go get`:

```
cd $GOPATH
go get github.com/lukemorton/api
cd src/github.com/lukemorton/api
go get ./...
```

**Running app:**

```
make
```

This will boot the app on [http://localhost:3000](http://localhost:3000)

## Deploying

**Deploy to now.sh:**

```
make deploy
```
