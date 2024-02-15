# GETH stack - Go, Echo, Turso and HTMX  by Quteo

Full text active search on the Edge with turso + fly.io. 


https://github.com/hanshoi/go-turso/assets/376152/656a20b3-e73f-4890-8d3f-e771b91977fc


Read my full thoughts from https://quteo.com/blog/replace-postgresql-with-2000-sqlite-databases

Uses Golang and the following technologies.

- HTMX
- Echo
- Templ
- Tailwind
- AlpineJS
- Turso

Graciously developed by https://quteo.com

As there are a bunch of things that need to be generated and built for a full server restart, we use Air to handle all that for us.

## Prequisities

- install turso-cli https://docs.turso.tech/cli/introduction
- install flyctl https://fly.io/docs/hands-on/install-flyctl/

```shell
npm install -D tailwindcss
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```

## Turso Remote DB

### Setup

Save token and db url to .env file
```shell
turso db create --enable-extensions test
turso token create test
turso db show test
```
After this you should have a .env file in root of this project that looks like this.

```shell
DB_URL=libsql://<db-name>-<username>.turso.io
DB_TOKEN=<some long token>
```
Now you are ready to roll.

### Usage locally

```shell
turso db shell solid-spitfire < create_tables.sql
air server --port 3000
```

## Deploy

You have a turso db already but to deploy it fully you need fly.io and their flyctl tool. 
```shell
fly launch
```
Check which url fly launched at and go there, this is your app now.



## Disclaimer

This project is for example and learning purposes. Go is not my main language, I'm just learning it while doing this.

So pardon my spaghetti code!

## License

Do whatever, I don't care.
