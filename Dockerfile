FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app ./src/main.go

ENV GIN_MODE=release
EXPOSE 80
CMD [ "app" ]