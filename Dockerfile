FROM node:20 as node-build

RUN mkdir -p /build
WORKDIR /build

COPY . .

RUN npx tailwindcss -i ./main.css -o ./static/tailwind.css

FROM golang:1.21-alpine as go-build

WORKDIR /usr/src/app

RUN go install github.com/a-h/templ/cmd/templ@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN templ generate
RUN go build -v -o /usr/local/bin/app .


FROM golang:1.21-alpine as run

RUN mkdir -p /web
WORKDIR /web
COPY --from=go-build /usr/local/bin/app /web/app
COPY --from=go-build /usr/src/app/.env /web/.env
COPY --from=node-build /build/static /web/static

CMD ["/web/app"]
