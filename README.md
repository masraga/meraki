# MERAKI

---

### About Meraki

Meraki is a mini framework with [go-fiber](https://docs.gofiber.io/) as core system. Meraki give a lot convenience for build RESTful api with [go-fiber](https://docs.gofiber.io/) and [MongoDB](https://www.mongodb.com/), cause meraki have a macro which is will be help developer to build model, repository and controller without write it manually, developer just need to specify name of a model or controller in command-line, then meraki will help to write default method and function automatically.

### Meraki Capabilities

- create automatic model, repository and controller
- only support mongodb driver database (will be updated)
- have default authentication system such as login, register, and validate token with JWT

### System Requirement

- Go 1.13 or above
- MongoDB 5.0 or above

### Installation

1. make sure `GOPATH` is defined in your machine environment
2. clone & open meraki project, then run:
   - `go get`
   - `go install`
   - `go build`
3. open terminal and run `meraki version` for makesure meraki can be use
4. run `make serve` in terminal and you will see meraki listen and serve http request

### Directory description

Below is directory description based on its utility
| Directory | Description |
| :-- | :-- |
| models | directory for define table / collection field for your database |
| controllers | bridge to connect a router to business logic |
| repositories | as a tools to communicate between business logic and database |
| usecase | all of business logic will be here |
| routes | as a tools for communicate between meraki and client side |
| middlewares | usually use for validating request or manipulate request before sent to business logic |
| server | main directory for serving and listening request |

### HOW TO USE COMMAND

all of commands will be start with `meraki` prefixes

#### creating model

model will be use to communicate with our database. everytime we create model, meraki will create model file in `models` directory and repository file in `repositories` directory. command to create model is :

```bash
meraki model --name ModelName
```

#### creating controller

controller will be use as a bridge between client and server. everytime we create controller, meraki will create controller file in `controllers` directory. command to create model is

```bash
meraki controller --name ModelName
```

Note: when creating `models` and `controllers`, make sure name of file starting with capital letter and not use special character except underscore symbol, e.g User, Store, Etc
