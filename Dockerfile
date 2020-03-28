# First, build the UI
FROM node:latest
ADD ui /src
WORKDIR /src
RUN npm install
RUN npm run build

# Secondly, compile the Go binary
FROM golang:latest
ENV CGO_ENABLED=0
ADD . /src
COPY --from=0 /src/build /src/ui/build
WORKDIR /src
RUN go generate
RUN go build

# Lastly, create a container with the resulting binary
FROM scratch
COPY --from=1 /src/guessing-game /usr/local/bin
ENTRYPOINT ["/usr/local/bin/guessing-game"]
