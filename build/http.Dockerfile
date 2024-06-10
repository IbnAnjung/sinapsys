# BUILD STAGE ############################################################
FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app

# Copy dependecies
COPY ./go.mod ./go.sum ./

# Download dependencies
RUN go mod download -x

# Copy source code
COPY . .

# Build the Go app
RUN go build -o main .

# RUN STAGE ############################################################
FROM alpine:3.20 AS runner

# Install dependencies when container running
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

# Set working directory
WORKDIR /app

# Copy main app from builder stage
COPY --from=builder /app/main .

EXPOSE 8000

# Command to run the executable
CMD ["./main"]