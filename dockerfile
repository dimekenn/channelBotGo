FROM golang:1.16 as builder
ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o telegramBotBeton -mod vendor .

FROM alpine:3.10
COPY --from=builder /build/telegramBotBeton /usr/local/bin
RUN chmod a+x /usr/local/bin/telegramBotBeton

CMD ["telegramBotBeton"]

