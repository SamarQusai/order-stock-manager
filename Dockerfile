FROM golang:1.22.0

# set working directory
WORKDIR $GOPATH/src/simple-order-stock-manager

# Copy the source code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o simple-order-stock-manager .
RUN go mod vendor

#EXPOSE the port
EXPOSE 9000

# Run the executable
CMD ["./simple-order-stock-manager"]
