FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app

ADD . /app

# Move to working directory /app
WORKDIR /app

# Copy and download dependency using go mod
#RUN go mod download

# Build the application
RUN go build -o /app/bin/server /app/cmd/server/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app/bin

# Export necessary port
EXPOSE 1337

# Command to run when starting the container
CMD ["/app/bin/server", "start", "--config", "prod.config"]