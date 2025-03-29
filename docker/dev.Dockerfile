# Use the official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@v1.42.0

# Copy go.mod and go.sum first to leverage Docker's caching mechanism
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (excluding files in .dockerignore)
COPY . .

# Expose the application port
EXPOSE 8000

# Set Air as the default command for hot-reloading development
CMD [ "air" ]
