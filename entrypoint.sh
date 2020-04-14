#!/bin/bash

set -e

until mysqladmin ping -h ${DB_HOST} -P ${DB_PORT} --silent; do
  echo "waiting for mysql..."
  sleep 2
done
echo "success to connect mysql."

# $GOPATH/bin/goose up
echo "migrated."

$GOPATH/bin/realize start
