FROM golang:1.23-alpine as base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

FROM base as dev
RUN go install github.com/air-verse/air@latest
COPY . .
EXPOSE 80
CMD ["sh", "-c", "air"]
