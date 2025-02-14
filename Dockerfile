FROM golang:1.22 as appbuilder

WORKDIR /app

COPY go.mod go.sum ./

# Download Go modules (dependencies)
RUN go mod download

# Copy the rest of the application code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/tenant-api ./api/main.go

FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the build stage
COPY --from=appbuilder /app/tenant-api .

# Expose the API port
EXPOSE 3000

# Run the binary
ENTRYPOINT ["./tenant-api"]