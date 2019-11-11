FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-service-vessel

COPY . .

RUN GOPROXY="https://goproxy.cn" GO111MODULE=on go mod download
RUN CGO_ENABLED=0 GOOS=linux GOPROXY="https://goproxy.cn" GO111MODULE=on go build -a -installsuffix cgo -o shippy-service-vessel

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-vessel .

CMD ["./shippy-service-vessel"]
