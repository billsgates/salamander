[![Build Status](https://travis-ci.com/billsgates/salamander.svg?branch=master)](https://travis-ci.com/github/billsgates/salamander)
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

### Current APIs

|             | Endpoint                    | Method | Status      | Description                               |
|-------------|-----------------------------|--------|-------------|-------------------------------------------|
| auth        | /auth/signup                | POST   | Complete    | User signup                               |
|             | /auth/signin                | POST   | Complete    | User login                                |
| users       | /users                      | GET    | Complete    | Get all users                             |
|             | /users                      | PATCH  | Complete    | Update user info                          |
|             | /users/{user_id}            | GET    | Complete    | Get user detail by user_id                |
| rooms       | /rooms                      | POST   | Complete    | Create a room                             |
|             | /rooms                      | GET    | Complete    | Get all joined rooms                      |
|             | /rooms/{room_id}            | GET    | Complete    | Get room detail by room_id                |
|             | /rooms/{room_id}            | PATCH  | Complete    | Update room by room_id                    |
|             | /rooms/{room_id}            | DELETE | Complete    | Delete room by room_id                    |
|             | /rooms/{room_id}/members    | GET    | To Do       | Get all members                           |
|             | /rooms/{room_id}/round      | POST   | Complete    | Add new round of the room                 |
|             | /rooms/{room_id}/round      | DELETE | To Do       | Remove current round of the room          |
|             | /rooms/{room_id}/invitation | POST   | Complete    | Create an invitation code in room room_id |
|             | /rooms/{room_id}/invitations| GET    | To Do       | Get all invitation codes in room room_id  |
|             | /rooms/join                 | POST   | Complete    | Join room by invitation code              |
|             | /rooms/join/{code}          | POST   | Complete    | Join room by url                          |
| services    | /services                   | GET    | Complete    | Get all services                          |
|             | /services/{service_id}      | GET    | Complete    | Get service detail by service_id          |
| participant | /participant                | DELETE | Complete    | Delete participant from room              |
|             | /participant/status         | PATCH  | To Do       | Change participant payment status         |
