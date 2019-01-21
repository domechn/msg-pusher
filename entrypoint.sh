#!/usr/bin/env bash


start_sender(){
    /app/sendmsg/sender start &
}

start_receiver(){
    /app/sendmsg/receiver start --addr-jaeger=$JAEGER_ADDR --addr-monitor=$MONITOR_ADDR
}

start_sender
sleep 1
start_receiver
