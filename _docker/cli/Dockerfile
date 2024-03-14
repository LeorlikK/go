FROM golang:1.21.5-alpine

WORKDIR /app

COPY ../../ .

RUN go mod download
#RUN go mod download  \
#    && go build -o main