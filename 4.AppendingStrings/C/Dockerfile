# Use an official C development image as the base image
FROM gcc:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the C source code to the working directory
COPY . .

# Build the C program
RUN gcc -g -o main main.c
RUN objdump -d main > disassembly.txt
RUN wc -l disassembly.txt

# Set the command to run when the container starts
CMD ["./main"]
