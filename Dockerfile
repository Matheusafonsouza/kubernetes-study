FROM golang:1.15

WORKDIR /app

COPY . .

RUN go build -o server .

EXPOSE 80

CMD ["./server"]