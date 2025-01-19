FROM golang:1.23-alpine

# Install required packages including sudo
RUN apk add --no-cache gcc musl-dev sudo

# Create user with specified UID/GID
ARG DOCKERUID
ARG DOCKERGID
RUN addgroup -g ${DOCKERGID} appuser && \
    adduser -D -u ${DOCKERUID} -G appuser appuser && \
    echo "appuser ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/appuser

WORKDIR /app

# Set ownership of the working directory
RUN chown -R appuser:appuser /app

# Copy go mod files
COPY backend/go.mod .
COPY backend/go.sum .

# Download dependencies
RUN go mod download

# Copy source code
COPY backend/ .

# Build the application
RUN go build -o main ./cmd/server

# Set ownership of all copied files
RUN chown -R appuser:appuser /app

# Switch to the new user
USER appuser

EXPOSE 8080

CMD ["./main"]
