## Определение базового образа для сборки
#FROM golang:latest as builder
#
## Установка рабочей директории внутри контейнера
#WORKDIR /app
#
## Копирование файлов проекта в контейнер
#COPY . .
#
## Сборка приложения
#RUN go build -o main ./cmd
#
## -------------
#
######### Start a new stage from scratch #######
#FROM alpine:latest
#
#WORKDIR /root/
#
## Copy the Pre-built binary file from the previous stage
#COPY --from=builder /app/main .
#
#RUN chmod +x main
#
#ENV PORT=:3333
#ENV LOG_LEVEL=debug
#
## Command to run the executable
#CMD ["./main"]

# Start from the latest golang base image
FROM golang:latest as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

ENV PORT=:1111
ENV LOG_LEVEL=Info
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=sasha
ENV POSTGRES_PASSWORD=12345678
ENV POSTGRES_DB=auction

# Command to run the executable
CMD ["./main"]