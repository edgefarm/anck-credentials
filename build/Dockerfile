FROM golang:1.16 AS builder
WORKDIR /code
COPY . .
RUN make

FROM alpine:3.14
COPY --from=builder /code/bin/anck-credentials /anck-credentials
ENTRYPOINT ["/anck-credentials"]
