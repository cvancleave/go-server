FROM golang:1.21 AS builder

ADD . /tmp/go-server
WORKDIR /tmp/go-server

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/server/main.go

FROM scratch
WORKDIR /root/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp/go-server/app /server

ENTRYPOINT ["/server"]
EXPOSE 4000
