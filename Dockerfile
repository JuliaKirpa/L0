FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./pg ./Consumer/cmd/main.go

FROM scratch
COPY --from=builder /app/pg /usr/bin/pg


ENTRYPOINT ["/usr/bin/pg"]