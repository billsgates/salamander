[![Build Status](https://travis-ci.com/billsgates/salamander.svg?branch=master)](https://travis-ci.com/github/billsgates/salamander)
# Salamander
backend api server for bills gate

## Development

I'v tried to follow the article's (https://ithelp.ithome.com.tw/articles/10241479) guide, take a look

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

### Requirements

* docker
* docker-compose
* git-crypt

### Quick Start

1. Decrypt the repository
```
$ git-crypt unlock path/to/key
```

2. Run the server
```
$ make run
$ docker ps -a
```

There should be six containers running in the background (`nginx`, `backend`, `mysql`, `mysql-adminer`, `swagger-ui`, `rabbitmq`), you can verify it by executing

* for `mysql-adminer`: visit (http://localhost:80/mysql-adminer) (server: `mysql`, username: `root`, password: `hermitcrab5566`, database: `hermitcrab_db`)
* for `swagger-ui`: visit (http://localhost:80/swagger-ui)
* for `rabbitmq`: visit (http://localhost:80/rabbitmq) (username: `guest`, password: `guest`)

3. Stop the server
```
$ make stop
```

### Current APIs

|             | Endpoint                    | Method | Status      | Description                               |
|-------------|-----------------------------|--------|-------------|-------------------------------------------|
| auth        | /auth/signup                | POST   | Complete    | User signup                               |
|             | /auth/signin                | POST   | Complete    | User login                                |
| users       | /user                       | GET    | Complete    | Get user info                             |
|             | /user                       | PATCH  | Complete    | Update user info                          |
|             | /user/{user_id}/rating      | GET    | Complete    | Get user Rating                           |
|             | /user/{user_id}/rating      | PATCH  | Complete    | Update user Rating                        |
| rooms       | /rooms                      | POST   | Complete    | Create a room                             |
|             | /rooms                      | GET    | Complete    | Get all joined rooms                      |
|             | /rooms/public               | GET    | Complete    | Get all public rooms                      |
|             | /rooms/{room_id}            | GET    | Complete    | Get room detail by room_id                |
|             | /rooms/{room_id}            | PATCH  | Complete    | Update room by room_id                    |
|             | /rooms/{room_id}            | DELETE | Complete    | Delete room by room_id                    |
|             | /rooms/{room_id}/start      | POST   | Complete    | Start room by room_id                     |
|             | /rooms/{room_id}/members    | GET    | Complete    | Get all members                           |
|             | /rooms/{room_id}/round      | POST   | Complete    | Add new round of the room                 |
|             | /rooms/{room_id}/round      | DELETE | Complete    | Remove current round of the room          |
|             | /rooms/{room_id}/invitation | POST   | Complete    | Create an invitation code in room room_id |
|             | /rooms/{room_id}/invitation | GET    | Complete    | Get all valid invitation codes in room    |
|             | /rooms/{room_id}/application| POST   | Complete    | Create an application to join room        |
|             | /rooms/{room_id}/application| GET    | Complete    | Get all applications of the room          |
|             | /rooms/join                 | POST   | Complete    | Join room by invitation code              |
|             | /rooms/join/{code}          | POST   | Complete    | Join room by url                          |
| services    | /services                   | GET    | Complete    | Get all services                          |
| participant | /participant                | DELETE | Complete    | Delete participant from room              |
|             | /participant/status         | PATCH  | Complete    | Change participant payment status         |
| application | /application/accept         | POST   | Complete    | Accept join room application              |
|             | /application/delete         | DELETE | Complete    | Delete join room application              |
