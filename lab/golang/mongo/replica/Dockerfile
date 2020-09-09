FROM golang:1.13.1-alpine3.10

WORKDIR /app

COPY main.go .

ENV GOPROXY "https://goproxy.io"

RUN env | grep GO

RUN go mod init 'github.com/silverswords/mongo-insight'
RUN go mod tidy
RUN go mod download

RUN go build -o main .

CMD ["./main"]