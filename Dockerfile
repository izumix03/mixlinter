FROM golang:1.12 as builder

WORKDIR /go/src/github.com/izumix03/mixlinter
COPY . .
RUN go get -u golang.org/x/tools/go/analysis
RUN CGO_ENABLED=0 go build -i -o /mixlinter -ldflags "-s -w" ./cmd/mixlinter

FROM golang:1.12-alpine

RUN apk add build-base
COPY --from=builder /mixlinter /

ENTRYPOINT ["/mixlinter"]
