FROM golang:1.21.4

WORKDIR /app
COPY . /app

RUN go mod download
RUN go build -o bin

EXPOSE 5600

ENTRYPOINT /app/bin