FROM golang:latest

USER root

# Install golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.0

# Install gofumpt
RUN go install mvdan.cc/gofumpt@latest

WORKDIR /app
COPY . .

CMD ["golangci-lint", "run", "--config", "/app/.golangci.yml"]
