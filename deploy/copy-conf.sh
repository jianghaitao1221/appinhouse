#!/bin/bash

APPINHOUSE_HOME=/srv/appinhouse
REDIS_CONFIG_PATH =/srv/redis/conf

sudo cp /tmp/appinhouse/docker-compose.yml $APPINHOUSE_HOME
sudo cp /tmp/appinhouse/redis.conf $REDIS_CONFIG_PATH