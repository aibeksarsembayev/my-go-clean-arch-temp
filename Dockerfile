FROM golang:1.17

WORKDIR /app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build -o build ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]