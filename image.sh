#!/usr/bin/env bash

su=""
if [[ $(id -u) = 0 ]]; then
    su="sudo"
fi

command="$1"
image_tag="$2"

args=""
for (( i=3; i<=$#; i++ )) do
    args+="${!i} "
done

if [[ "${command}" = "build" ]]; then
    shcmd="docker build -t ${image_tag} ."
elif [[ "${command}" = "run" ]]; then
    shcmd="docker run ${args} ${image_tag}"
elif [[ "$command" = "push" ]]; then
    echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin \
        && shcmd="docker push ${image_tag}"
elif [[ "${command}" = "remove" ]]; then
    shcmd="docker rmi ${image_tag}"
else
  echo "command expected 'build/push/remove'"
  exit -1
fi

echo "${su} ${shcmd}" && sh -c "${shcmd}"
