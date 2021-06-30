FROM golang:1.16.5-buster
ENV CGO_ENABLED=0
RUN go install github.com/spf13/cobra/cobra@latest
