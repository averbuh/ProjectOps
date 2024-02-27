FROM golang:1.21.1-alpine

# Set destination for COPY
WORKDIR /app




# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["cmd/standart"]