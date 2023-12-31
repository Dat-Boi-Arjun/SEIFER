# syntax=docker/dockerfile:1
ARG fp=/root

FROM golang:latest as build_stage
ARG fp

COPY ./dispatcher_pod/config_step/deploy_container/ ${fp}/src/dispatcher_pod/config_step/deploy_container/

ENV GOPATH=$fp
# Initialize module in src directory (root of code files)
WORKDIR ${fp}/src
RUN go mod init github.com/Dat-Boi-Arjun/DEFER
RUN go get k8s.io/api@latest
RUN go get k8s.io/client-go@latest
RUN go get k8s.io/apimachinery@latest
RUN go mod tidy

WORKDIR ${fp}/src/dispatcher_pod/config_step/deploy_container/bin/
RUN env GOOS=linux GOARCH=arm go build -o ${fp}/main main.go

# Slimmed build stage only w/ executable
FROM alpine:latest
ARG fp

WORKDIR $fp
COPY --from=build_stage ${fp}/main ./

# Install kubectl
RUN apk update && \
      apk add curl && \
      curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/arm64/kubectl && \
      chmod +x ./kubectl && \
      mv ./kubectl /usr/local/bin/kubectl

ENTRYPOINT ["./main"]