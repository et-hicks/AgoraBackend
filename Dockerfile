# Use base golang image from Docker Hub
FROM golang:1.16 AS build

WORKDIR /hello-world

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app.
# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN echo "Go gcflags: ${SKAFFOLD_GO_GCFLAGS}"
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -mod=readonly -v -o /app
RUN echo "I guess google does use the dockerfile. neat. I wonder what this thing does"
RUN echo "oh well. I just copy things from the internet. it mostly works out ok"
# Now create separate deployment image
FROM gcr.io/distroless/base

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

# Copy template & assets
WORKDIR /hello-world
COPY --from=build /app ./app
#COPY index.html index.html
#COPY assets assets/

ENTRYPOINT ["./app"]