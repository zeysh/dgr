#!/dgr/bin/busybox sh
set -e
. /dgr/builder/export
. /dgr/bin/functions.sh
isLevelEnabled "debug" && set -x

onError() {
    if [ "${CATCH_ON_ERROR}" == "true" ]; then
        echo_red "${1} failed. dropping to shell in build"
        sh
    fi
    exit 1
}

execute_files "/dgr/runlevels/inherit-build-early" || onError "Inherit-build-early"
execute_files "/dgr/builder/runlevels/build" || onError "Build"

if [ "${CATCH_ON_STEP}" == "true" ]; then
    echo_purple "Catch requested dropping to shell in build"
    sh
fi
