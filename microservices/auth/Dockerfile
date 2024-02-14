# Development stage
FROM golang:latest AS development

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Install air for live reloading in development
RUN go get -u github.com/cosmtrek/air

# Production stage
FROM golang:latest AS production

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application without air for production
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final stage
FROM alpine:latest AS final

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary built from the production stage
COPY --from=production /app/app .

# Expose the port on which your application will run
EXPOSE 8080

# Command to run the executable
CMD ["./app"]