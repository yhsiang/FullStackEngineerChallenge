#!/bin/sh

set -e

host="$1"

until mysqladmin ping -h "$host" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD"; do
  >&2 echo "MySQL is unavailable - sleeping"
  sleep 1
done

>&2 echo "MySQL is up - executing command"
/go/bin/goose status
/go/bin/goose up
