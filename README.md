# Salamander
backend api server for our project

## Development

I'v tried to follow the article's (https://ithelp.ithome.com.tw/articles/10241479) guide, take a look

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

### Requirements

* docker
* docker-compose

### Quick Start

1. Run the server
```
$ make run
$ docker ps -a
```

There should be four containers running in the background (`backend`, `mysql`, `mysql-adminer`, `swagger-ui`), you can verify it by executing

* for `mysql-adminer`: visit (http://localhost:8080) (server: `mysql`, username: `root`, password: `hermitcrab5566`, database: `hermitcrab_db`)
* for `swagger-ui`: visit (http://localhost:8081)

2. Stop the server
```
$ make stop
```
