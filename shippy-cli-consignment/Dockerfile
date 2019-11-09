FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/shippy-cli-consignment

COPY . .

RUN GOPROXY="https://goproxy.cn" GO111MODULE=on go mod download

RUN CGO_ENABLED=0 GOOS=linux GOPROXY="https://goproxy.cn" GO111MODULE=on go build -a -installsuffix cgo -o shippy-cli-consignment

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/shippy-cli-consignment/shippy-cli-consignment .
COPY --from=builder /app/shippy-cli-consignment/consignment.json .

CMD ["./shippy-cli-consignment"]
