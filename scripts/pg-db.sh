#!/usr/bin/env bash

#!/usr/bin/env bash

set -e

scripts_dir=$(realpath "$(dirname $0)")
parent_dir=$(realpath "$(dirname $0)/..")
source $parent_dir/.env
source $scripts_dir/config.sh

echo ${POSTGRES_PORT}

runContainer() {
  pRun docker run -it -d \
    --name $db_container_name \
    -p ${POSTGRES_PORT}:${POSTGRES_DOCKER_PORT} \
    --restart unless-stopped \
    -v ./ezd_daemon_pg_db_data:/var/lib/postgresql \
    -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
    -e POSTGRES_USER=${POSTGRES_USER} \
    -e POSTGRES_DB=${POSTGRES_DB} \
    $db_image_name
}
stopContainer() {
  pRun docker stop $db_container_name
  pRun docker rm $db_container_name
}
buildImage() {
  local commit=$(git rev-parse HEAD | head -c 8)
  local timestamp=$(date -u +"%Y%m%dT%H%M%SZ")
  local timestamp_iso=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  local tags=("$timestamp.$commit" "latest")
  local tagStr=""
  for i in ${tags[@]}; do
    # echo $i
    tagStr="$tagStr $db_image_name:$i"
  done
  echo ${tags[@]}
  echo ${tagStr}
  if [[ $(basename $parent_dir) != $PROJ_DIRNAME ]]; then
    fatal "wrong directory"
  fi
  echo "${parent_dir}/Dockerfile.pg"
  docker build --pull --no-cache \
    -t "${db_image_name}:${commit}" \
    -f "${parent_dir}/Dockerfile.pg" .
  for image_tag in ${tags[@]}; do
    docker image tag ${db_image_name}:${commit} ${db_image_name}:${image_tag}
  done
}

main() {
  local arg1=$1
  case "$arg1" in
    b|build)
      buildImage
      ;;
    s|stop)
      stopContainer
      ;;
    *) if [ -z "$arg1" ]; then
        runContainer
      else
        fatal "Invalid argument: $arg1"
      fi
      ;;
  esac
}

main "$@"

