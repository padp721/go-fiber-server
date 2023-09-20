FROM golang:1.21.1-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main .
CMD [ "/app/main" ]