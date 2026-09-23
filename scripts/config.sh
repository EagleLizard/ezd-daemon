#!/usr/bin/env bash


PROJ_DIRNAME=ezd-daemon

db_image_name=ezd_daemon_pg_db
db_container_name=ezd-daemon-postgres

# scripts_dir=$(realpath "$(dirname $0)")

echo $(date -u +"%Y-%m-%dT%H:%M:%SZ")

fatal() {
	echo '[FATAL]' "$@" >&2
	exit 1
}
pRun() {
  echo "_> $@"
  eval $@
}
