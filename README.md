# Go Microframework
   
Go Microframework built on top of [Uber FX](https://go.uber.org/fx) and various Golang packages, fully supporting [Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection), [CQRS](https://microservices.io/patterns/data/cqrs.html) && [Event Sourcing](https://microservices.io/patterns/data/event-sourcing.html).

Easy to build ([Domain Driven Design](https://en.wikipedia.org/wiki/Domain-driven_design) && [Event Driven](https://en.wikipedia.org/wiki/Event-driven_architecture)) applications.

## Packages

[Uber FX](https://go.uber.org/fx) - A dependency injection framework.

[Viper](https://github.com/spf13/viper) - A complete configuration solution for Go applications including [12-Factors](https://12factor.net/) app.

[Echo](https://echo.labstack.com/) - Handle && manage HTTP requests.

[GORM](http://gorm.io/) - Handle && manage database connections, ORM library.

[Watermill](https://watermill.io/) - A message streams library, be used for CQRS / Event Sourcing / RPC over messages / Sagas.