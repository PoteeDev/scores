FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY handlers handlers
COPY models models
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /scores .

FROM alpine
WORKDIR /usr/app
COPY --from=build /scores .
ENV PORT=8080
ENTRYPOINT [ "./scores"]
