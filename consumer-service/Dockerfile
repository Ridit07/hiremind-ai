# Stage 1: Builder
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata openssh-client

WORKDIR /app

# Configure Git for private repos
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Copy SSH key for private repo access (build arg)
ARG SSH_PRIVATE_KEY
RUN mkdir -p /root/.ssh && \
    echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan -H github.com >> /root/.ssh/known_hosts

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer-service .

# Stage 2: Runtime
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /home/appuser

COPY --from=builder --chown=appuser:appuser /app/consumer-service .

USER appuser

EXPOSE 50051

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD nc -z localhost 50051 || exit 1

CMD ["./consumer-service"]