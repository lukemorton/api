# \*.api.lukemorton.co.uk

## Prerequisites

- Install [go](https://golang.org/dl/) :D **required**
- Install [now-cli](https://github.com/zeit/now-cli/) if you're deploying to now
- Install [docker](https://www.docker.com/products/overview#/install_the_platform) if you're testing the docker file locally

## Development

Setup hosts:

```
sudo echo "users  127.0.0.1" >> /etc/hosts
sudo echo "authors  127.0.0.1" >> /etc/hosts
```

Run docker compose via Makefile:

```
make
```

Now visit any of the following:

 - [https://localhost](https://localhost)
 - [https://users](https://users)
 - [https://authors](https://authors)
