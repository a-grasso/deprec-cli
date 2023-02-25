FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY config.json ./

RUN go build -o /deprec-cli .

FROM alpine

WORKDIR /app

COPY --from=build /deprec-cli /deprec-cli

# chmod
# alias /deprec-cli deprec-cli