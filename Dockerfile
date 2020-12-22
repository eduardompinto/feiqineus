FROM golang:alpine
WORKDIR /app
COPY *.go /app
COPY go.mod /app
COPY go.sum /app
COPY internal /app/internal
EXPOSE 8080
CMD ["go", "run", "."]