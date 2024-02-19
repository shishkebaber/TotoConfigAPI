FROM golang:1.21


WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# Copy IPs database. 
#TODO: better to store it somewhere else, some Cloud storage probably
COPY *.BIN ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /toto-config-api ./cmd/main.go


EXPOSE 8080

# Run
CMD ["/toto-config-api"]