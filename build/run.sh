#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/speedtest" -importPath speedtest -srcPath "$SCRIPTPATH/src" -runMode dev
