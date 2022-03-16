FROM golang:1.7.1

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/DanielPickens/Obsidian

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .


# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080
EXPOSE 51550

# Run the executable
CMD ["go-sample-app"]
