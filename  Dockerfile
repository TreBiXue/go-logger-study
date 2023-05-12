FROM golang:1.17
COPY . /app
WORKDIR /app
RUN go build -o app
CMD ["/app/app"]