FROM golang:1.23

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o fullcycle-clean-arch .

CMD ["./fullcycle-clean-arch"]
