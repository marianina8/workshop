FROM alpine:latest

RUN apk add --no-cache go

COPY . .

CMD ["go", "test", "-v", "github.com/marianina8/example30/cmd"]