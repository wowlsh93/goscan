#!/bin/bash

COMMAND_goscan="../bin/bs scanner start"

case "$1" in
  start)
    echo "Starting goscan.."
    nohup $COMMAND_goscan &
    echo "Done!!"
    ;;
  stop)
    echo "Stopping goscan.."
    pkill -9 goscan
    echo "Done!!"
    ;;
  *)
    echo "Usage: ./goscan.sh {start|stop}"
    exit 1
    ;;
esac

exit 0
