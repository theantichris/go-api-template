# go-api-template

![Go](https://github.com/theantichris/go-api-template/workflows/Go/badge.svg)

A template repo for an API server written in Go.

## Dependencies

* [Godotenv](https://github.com/joho/godotenv)
* [Mux](https://github.com/gorilla/mux)
* [Logrus](https://github.com/sirupsen/logrus)
* [Sentry](https://github.com/getsentry/sentry-go) - optional

## Quick start

1. Copy `.sample.env` to `.env` and update values
1. Run `go build`
1. Run `./go-api-template`

## Deployment

A Procfile is included to easy deployment to [Heroku](https://www.heroku.com/)

## Endpoints

### GET /health

Response

```json
{
  "alive": true
}
```
