FROM golang:latest as builder

# Install Go Dep
RUN go get github.com/golang/dep/cmd/dep

# Navigate to project dir
WORKDIR /go/src/eman/passport/game

# Install dependencies
COPY Gopkg.toml .
COPY Gopkg.lock .
RUN dep ensure -vendor-only

# Copy source code
COPY ./src ./src
COPY ./res ./res

# Build and Run
RUN env GOOS=linux go build -ldflags="-s -w" -o bin/main src/server.go
# CMD ["./bin/main"]

FROM scratch
COPY --from=builder  /go/src/eman/passport/game/bin /app

#CMD ["./app/main"]

#EXPOSE 8181