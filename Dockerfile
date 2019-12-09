# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang

# Add Maintainer Info
LABEL maintainer="Thomas Pickett <test@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/go-test-rest-api

# COPY go.mod .
# COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN go get ./

# Build the Go app
RUN go build -o go-test-rest-api .

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["./go-test-rest-api"]