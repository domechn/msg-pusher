FROM alpine:latest

RUN mkdir -p /app/sendmsg/conf

ADD ./dist/* /app/sendmsg/
ADD ./conf.yaml /app/sendmsg/conf
ADD ./entrypoint.sh /app/sendmsg

RUN chmod +x /app/sendmsg/entrypoint.sh

EXPOSE 8990

ENTRYPOINT ["/bin/sh","/app/sendmsg/entrypoint.sh"]