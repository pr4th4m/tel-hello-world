# # Build image
# docker build -t hello-world .
#
# # Run container
# docker run --rm -p 3000:3333 hello-world
# docker logs -f hello-world
FROM golang:1.10.0-stretch

# Copy binary and config file
WORKDIR /go/src/github.com/pratz/tel-hello-world
COPY . .

# Run unit tests
RUN go test -v

# Build binary
RUN go build

# Run app
EXPOSE 3333
ENTRYPOINT ["./tel-hello-world"]
