FROM alpine:latest

RUN mkdir -p /app/sendmsg/conf

ADD ./dist/sender /app/sendmsg
ADD ./conf.yaml /app/sendmsg/conf
RUN apk add --no-cache tzdata

ENTRYPOINT ["/app/sendmsg/sender","start"]