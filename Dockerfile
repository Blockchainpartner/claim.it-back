FROM golang:1.12

# Copy sources
RUN mkdir /go/src/claim.it
COPY . /go/src/claim.it
WORKDIR /go/src/claim.it

# Run the project
CMD ["go", "run", "main.go"]
