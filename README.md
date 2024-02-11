# HETTA - Htmx + Echo + Templ + Tailwind + AlpineJS by Quteo

This is a example project for getting following technologies to work with golang.

- HTMX
- Echo
- Templ
- Tailwind
- AlpineJS
- Turso

Graciously developed by https://quteo.com

As there is a bunch of things that need to be generated and built for a full server restart, we use Air to handle all that for us.

## Prequisities

- install turso-cli https://docs.turso.tech/cli/introduction

```shell
npm install -D tailwindcss
go install github.com/cosmtrek/air@latest
go install -tags 'sqlite' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Run

```shell
turso dev --db-file my.db
migrate -database "sqlite://my.db" -path db/migrations up
air server --port 3000
```

## Turso Remote DB

### Setup

Save token and db url to .env file
```shell
turso db create --enable-extensions test
turso token create test
turso db show test
```

### Usage

```shell
turso db shell solid-spitfire < db/migrations/000001_initial.up.sql
air server --port 3000
```

