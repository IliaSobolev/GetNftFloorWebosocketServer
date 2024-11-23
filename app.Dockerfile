FROM golang:1.23

# Set app workdir
WORKDIR /go/src/app

# Copy dependencies list
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application sources
COPY . .

#Tests
RUN go test

# Build app
RUN go build -o app .

# Run app
CMD ["./app"]