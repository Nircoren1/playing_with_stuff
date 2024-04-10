# syntax=docker/dockerfile:1

FROM golang:1.19 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM build-stage AS run-test-stage

RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=builder /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

# Command to run the executable
ENTRYPOINT ["/docker-gs-ping"]
