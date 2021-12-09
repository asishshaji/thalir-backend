FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM scratch

ENV SERVER_PORT=9090
ENV HOST=db
ENV DB_NAME=thalir
ENV DB_USER=asish
ENV DB_PASSWORD=kai1253
ENV DB_PORT=5432
# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]