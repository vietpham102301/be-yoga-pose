FROM golang:latest
LABEL authors="vietpham1023"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o main .

CMD ["./main"]