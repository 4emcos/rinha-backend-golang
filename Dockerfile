FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go get -d -v ./...

RUN go build -o /usr/local/bin/app ./src/main.go

CMD [ "app" ]