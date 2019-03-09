FROM golang:1.12

# Copy sources
RUN mkdir -p /go/src/github.com/Blockchainpartner/claim.it-back
COPY . /go/src/github.com/Blockchainpartner/claim.it-back
WORKDIR /go/src/github.com/Blockchainpartner/claim.it-back

# Run the project
CMD ["go", "run", "main.go"]
