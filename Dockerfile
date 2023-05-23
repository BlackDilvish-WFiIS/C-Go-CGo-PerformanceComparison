# Use an official Golang runtime as the base image
FROM golang:latest

# Install build dependencies for C and CGo
RUN apt-get update && \
    apt-get install -y build-essential

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
# RUN go mod download

# Build the Go program
# RUN go build -gcflags "-N -l" -o main
# RUN go tool objdump -S main > disassembly.txt
# RUN wc -l disassembly.txt

# Set the command to run when the container starts
# CMD ["bash -c run.sh"]
