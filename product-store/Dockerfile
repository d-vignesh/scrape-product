FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-store *.go

FROM alpine:latest

WORKDIR /app

EXPOSE 9002

COPY --from=builder /build/product-store .

CMD ["./product-store"]