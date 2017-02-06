# \*.api.lukemorton.co.uk

A concert of micro service APIs.

## Prerequisites

- Install [go](https://golang.org/dl/) :D **required**
- Install [docker](https://www.docker.com/products/overview#/install_the_platform) if you're developing and the micro service APIs in concert

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

 - [http://localhost](http://localhost)
 - [http://users](http://users)
 - [http://authors](http://authors)
