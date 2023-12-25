FROM golang:1.21.4

WORKDIR /app
COPY . /app

RUN go mod download
RUN go build -o ./build/bin .

ENTRYPOINT /app/build/bin