FROM golang:1.21

WORKDIR /notifier

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /docker-bn-rest-app ./cmd/notifier

EXPOSE 8080

CMD ["/docker-bn-rest"]