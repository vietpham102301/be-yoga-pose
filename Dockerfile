# Use the official Golang image as the base image for the Go part
FROM golang:1.21.4 as go_builder

# Set the working directory for Go
WORKDIR /app

# Copy Go application files
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go application for AMD64
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o main .

# Use the official Python image as the base image for the Python part
FROM python:3.11.5-slim as python_builder

# Set the working directory for Python
WORKDIR /app

# Copy Python application files
COPY python_dependencies.txt .
RUN pip install -r python_dependencies.txt

# Copy the Go binary from the first stage
COPY --from=go_builder /app/main .

# Set the working directory for the combined image
WORKDIR /app

# Copy image files
COPY saved_frames/images/ saved_frames/images/
COPY saved_frames/cropped_images/ saved_frames/cropped_images/
COPY python python

# Expose port for the Go application
EXPOSE 8080

# Run the Go application
CMD ["./main"]
