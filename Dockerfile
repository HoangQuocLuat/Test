FROM golang:1.21.1

WORKDIR /app

COPY . .

RUN  go mod tidy

EXPOSE 8888

CMD ["go", "run", "main.go"]
