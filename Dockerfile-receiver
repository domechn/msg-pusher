FROM alpine:latest

RUN mkdir -p /app/sendmsg/conf

ADD ./dist/receiver /app/sendmsg
ADD ./conf.yaml /app/sendmsg/conf

EXPOSE 8990

ENTRYPOINT /app/sendmsg/receiver start --addr-jaeger=$JAEGER_ADDR --addr-monitor=$MONITOR_ADDR