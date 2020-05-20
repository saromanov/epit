FROM golang:alpine as builder

ADD ./backend/app /build/backend/app
ADD ./backend/go.mod /build/backend/go.mod
ADD ./backend/go.sum /build/backend/go.sum

RUN go mod download 
RUN go build -o app ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /src/goapp /app/
ENTRYPOINT ./goapp