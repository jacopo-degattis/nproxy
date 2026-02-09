FROM golang:1.25-alpine

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /nproxy
CMD ["/nproxy"]
