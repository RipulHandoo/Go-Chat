FROM golang:1.20-bullseye AS build-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only files required for go mod download
COPY go.mod go.sum ./
COPY .env ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN --mount=type=cache,target=/go/pkg/mod \
   --mount=type=cache,target=/root/.cache/go-build \
   go mod download

FROM build-dev AS build-production
# Copy the source FROM the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o api-golang

FROM scratch

WORKDIR /

# Copy the Pre-built binary file FROM the previous stage
COPY --from=build-production /app/api-golang api-golang

EXPOSE 8080

CMD [ "/api-golang" ]