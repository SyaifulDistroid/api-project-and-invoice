FROM golang:1.16-alpine
WORKDIR /EMI/
COPY go.mod go.sum ./
RUN go mod download
COPY COPY . .
COPY EMI/main.go .
WORKDIR /EMI
RUN go build -o main .
EXPOSE 8585
CMD ["./EMI/main"]
