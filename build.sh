#!/bin/bash
# resolve links - $0 may be a softlink
if [ -z "$goscan_HOME" ];then
  PRG="$0"
  while [ -h "$PRG" ] ; do
    ls=$(ls -ld "$PRG")
    link=$(expr "$ls" : '.*-> \(.*\)$')
    if expr "$link" : '/.*' > /dev/null; then
      PRG="$link"
    else
      PRG=$(dirname "$PRG")/"$link"
    fi
  done
  cd "$(dirname "$PRG")" || exit 1
  export goscan_HOME="$(pwd)"
  cd - &>/dev/null || exit 1
fi

echo $goscan_HOME

case $1 in
  "build" | "")
    cd $goscan_HOME/cmd/bs
    go build ;;
  "test")
    echo "test your application"
    ;;
  "dist")
    cd $goscan_HOME/cmd/bs
    go build
    cd -
    mkdir -p $goscan_HOME/assembly/dist/bin/
    mkdir -p $goscan_HOME/assembly/dist/logs/
    cp $goscan_HOME/cmd/bs/bs $goscan_HOME/assembly/dist/bin/
    tar -czf $goscan_HOME/assembly/goscan.tar.gz -C $goscan_HOME/assembly/dist .
    ;;
  *) echo "Unknown command: $1";;
esac