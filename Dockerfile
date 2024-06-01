FROM golang:1.22.3 as build

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o stress-tester .

ENTRYPOINT ["./stress-tester"]
