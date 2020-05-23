FROM golang:alpine as builder

RUN mkdir epit
ADD . /epit
WORKDIR /epit

RUN go mod download 
RUN go build -o epit ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /epit/epit /bin/epit
ENTRYPOINT epit