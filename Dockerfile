FROM golang:1.13 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

ADD . /workspace

WORKDIR /workspace

RUN go build -o app

FROM alpine

COPY --from=builder /workspace/app /app

ENTRYPOINT [ "/app" ]