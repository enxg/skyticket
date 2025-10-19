FROM golang:1.25-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY internal/ ./internal/
COPY pkg/ ./pkg/

RUN make docs
RUN make build-prod

EXPOSE 3000

CMD ["/app/skyticket"]
