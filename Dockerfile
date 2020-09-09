FROM golang:1.15.1 as builder

WORKDIR /app
COPY . .

RUN go get -u github.com/cosmtrek/air
RUN go build -o app main.go

FROM alpine

COPY --from=builder /app/app /app
EXPOSE 80
ENTRYPOINT ["/app"]
