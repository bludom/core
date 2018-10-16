#!/bin/bash

export ARCH=${1:-'ARM'}
docker-compose up -d
