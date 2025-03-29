# Stage 1: Build
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o 472-BE-CICD ./main.go

# Stage 2: Final Image
FROM alpine:latest

WORKDIR /app

# Install required dependencies and ca-certificate for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the build binary
COPY --from=builder /app/472-BE-CICD .

# Change access matrix to be executatble
RUN chmod +x 472-BE-CICD

# Copy timezone info
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip

# Set default port
EXPOSE 8080

CMD [ "./472-BE-CICD" ]
