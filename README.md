# GoFiber Docker Boilerplate

![Release](https://img.shields.io/github/release/gofiber/boilerplate.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/boilerplate/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/boilerplate/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/boilerplate/workflows/Linter/badge.svg)


## Development

### Add File .env 

Copy file from .env.example to create .env for your local work 

### Start the application from docker

```bash
cd build
docker-compose build --no-cache 
docker-compose up
docker run -d -p 3000:3000 gofiber
```

### Start the application from local

```bash
go run main.go
```

## Production

Not yet configuration

![Go Fiber Docker Boilerplate](./go_fiber_boilerplate.gif)

## Feature

Feature planning to embed for this boilerplate
- DDD(Domain Driven Design) architecture (complete)
- Auto Migration Db
- Jwt Auth Basic
- Request Validation (complete)
- Redis/Elastic Cache
- Logging System(With DB)
- Unit Testing
- Role Management
- Default Error Handler And Custom Error Handler