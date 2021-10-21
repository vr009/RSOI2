#!/usr/bin/env bash
set -e

export VARIANT=v4
export SCRIPT_PATH=/docker-entrypoint-initdb.d/
export PGPASSWORD=test
psql -U program -d services -f "$SCRIPT_PATH/schemes/schema-$VARIANT.sql"