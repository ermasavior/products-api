FROM golang:1.19-alpine as Build

COPY . .

RUN GOPATH= go build -o /main cmd/main.go

FROM alpine:latest

COPY --from=Build /main .

EXPOSE 1323

ENTRYPOINT ["./main"]