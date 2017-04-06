# \*.api.lukemorton.co.uk

A concert of micro service APIs.

## About Clean Architecture

Let's face it, there are a lot of design patterns and ways of structuring code outside of MVC. Many web developers these days know MVC has it's limits and choose to move their domain logic outside of the confines of it.

When you have the freedom of structuring your own code, sometimes you can end up in just as much mess. A lot of engineers will start by placing code in a services folder, or plain old objects in a lib folder. This gets you so far but without a way of organising said logic, things can still get out of control and are even less documented than your traditional MVC framework.

Clean Architecture is a term coined by Uncle Bob, and we at Made Tech have begun defining [our own flavour](https://github.com/madetech/clean-architecture).

This repository is an example of a group of Golang microservices that use Clean Architecture instead of using MVC. To be clear, you can easily use Clean Architecture with MVC. It just wasn't necessary here for this example.

There are three patterns that make up Clean Architecture and they are use cases, domain models, and gateways.

### Use cases

A use case represents a particular action performed by a user. It might be a view of latest orders, in which case you might call it `ViewLatestOrders`, or perhaps creating a new order `CreateOrder`. In OOP languages it would likely be a class with a single public method, in Golang I've used functions instead.

One example from this repository is that of the user `users.Register` use case.

- Use case: https://github.com/lukemorton/api/blob/master/users/register.go
- Test for use case: https://github.com/lukemorton/api/blob/master/users/register_test.go

### Domain models

Use cases will want to act on particular things. Our `users.Register` use case creates a `users.User` domain model. The primary purpose of this structure is to hold the data attributed to a particular user. It may also have some logic associated with it in the shape of methods. Usually only logic that is shared amongst multiple use cases and always performed on the domain models data will be kept in the domain model. Logic unrelated to a domain model would be best placed in child use cases.

An example of a domain logic is that of hashing and verifying a users password.

- Domain model: https://github.com/lukemorton/api/blob/master/users/user.go

### Gateways

Finally, use cases will want to cause side effects on the outside world. That might be calling an API, accessing a database, or saving something to a disk (heaven forbid). We use gateways so that when testing our use cases we can inject mock gateways to keep our tests as fast a possible. By defining interfaces for individual methods on the gateway it becomes super easy to move from say AWS DynamoDB to PostgreSQL if you decided you needed the features of a relational DB. In the Rails world, you would likely use ActiveRecord inside a gateway.

You test your gateway too and you would usually not want to mock this out. By keeping your side effects in a smaller surface area, they become quicker and easier to test.

An example of a gateway from this repository would be that of the `users.SQLUserStore`.

- Gateway: https://github.com/lukemorton/api/blob/master/users/sql_user_store.go
- Tests for Gateway: https://github.com/lukemorton/api/blob/master/users/sql_user_store_test.go
- Interfaces for Gateway: https://github.com/lukemorton/api/blob/master/users/user_store.go

### Find out more about Clean Architecture

https://github.com/madetech/clean-architecture

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
