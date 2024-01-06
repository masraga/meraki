# MERAKI
---

### About Meraki
Meraki is boilerplate for build awesome restful API with golang. Meraki use [gin-gonic](https://github.com/gin-gonic/gin) as core system. With default authentication handler, and command-line syntax for write code such as model, repositories, and controller, will help developer to focus build business logic and save time to deliver an application.

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
| :--       | :--         |
| models    | directory for define table / collection field for your database |
| controllers | bridge to connect a router to business logic |
| repositories | as a tools to communicate between business logic and database |
| usecase | all of business logic will be here |
| routes | as a tools for communicate between meraki and client side |
| middlewares | usually use for validating request or manipulate request before sent to business logic |
| server | main directory for serving and listening request |