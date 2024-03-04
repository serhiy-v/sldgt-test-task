FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd ./cmd && env CGO_ENABLED=0 GOOS=linux go build -o /sldgt-test-task

EXPOSE 8080

CMD ["/sldgt-test-task"]
