FROM golang:1.9-alpine AS build_base



# Set the Current Working Directory inside the container
WORKDIR /Obsidian

# We want to populate the module cache based on the go.{mod,sum} files.




COPY . .



# Build the Go app


# Start fresh from a smaller image
FROM alpine:3.9 



# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/app/Obsidian"]
