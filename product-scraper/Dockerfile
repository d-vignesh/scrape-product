From golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o product-scraper *.go

FROM alpine:latest

WORKDIR /app

EXPOSE 9001

COPY --from=builder /build/product-scraper .

CMD ["./product-scraper"]

