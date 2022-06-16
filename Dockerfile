# Thanks to https://github.com/evrone/go-clean-template
# Step 1: Modules caching
# Add go.sum later
FROM golang:1.18-alpine as modules
COPY go.mod /modules/ 
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.18-alpine as builder
# COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /bin/app /app
CMD ["/app"]