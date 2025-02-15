#!/bin/bash
# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

###
# Runs all Mainflux microservices (must be previously built and installed).
#
# Expects that PostgreSQL and needed messaging DB are alredy running.
# Additionally, MQTT microservice demands that Redis is up and running.
#
###

BUILD_DIR=../build

# Kill all mainflux-* stuff
function cleanup {
    pkill mainflux
    pkill nats
}

###
# NATS
###
nats-server &
counter=1
until fuser 4222/tcp 1>/dev/null 2>&1;
do
    sleep 0.5
    ((counter++))
    if [ ${counter} -gt 10 ]
    then
        echo "NATS failed to start in 5 sec, exiting"
        exit 1
    fi
    echo "Waiting for NATS server"
done

###
# Users
###
MF_USERS_LOG_LEVEL=info MF_USERS_HTTP_PORT=9002 MF_USERS_GRPC_PORT=7001 MF_USERS_ADMIN_EMAIL=admin@mainflux.com MF_USERS_ADMIN_PASSWORD=12345678 MF_EMAIL_TEMPLATE=../docker/templates/users.tmpl $BUILD_DIR/mainflux-users &

###
# Things
###
MF_THINGS_LOG_LEVEL=info MF_THINGS_HTTP_PORT=9000 MF_THINGS_AUTH_GRPC_PORT=7000 MF_THINGS_AUTH_HTTP_PORT=9002 $BUILD_DIR/mainflux-things &

###
# HTTP
###
MF_HTTP_ADAPTER_LOG_LEVEL=info MF_HTTP_ADAPTER_PORT=8008 MF_THINGS_AUTH_GRPC_URL=localhost:7000 $BUILD_DIR/mainflux-http &

###
# WS
###
MF_WS_ADAPTER_LOG_LEVEL=info MF_WS_ADAPTER_PORT=8190 MF_THINGS_AUTH_GRPC_URL=localhost:7000 $BUILD_DIR/mainflux-ws &

###
# MQTT
###
MF_MQTT_ADAPTER_LOG_LEVEL=info MF_THINGS_AUTH_GRPC_URL=localhost:7000 $BUILD_DIR/mainflux-mqtt &

###
# CoAP
###
MF_COAP_ADAPTER_LOG_LEVEL=info MF_COAP_ADAPTER_PORT=5683 MF_THINGS_AUTH_GRPC_URL=localhost:7000 $BUILD_DIR/mainflux-coap &

trap cleanup EXIT

while : ; do sleep 1 ; done
