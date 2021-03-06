#!/usr/bin/env bash


usage()
{
    echo "Usage: ${0##*/} {install|update}"
    exit 1
}

ACTION=$1

if [ -z "$ACTION" ]; then
   usage
fi



echo 'start'

cd ../../../..
CURDIR=`pwd`
HOME=env | grep ^HOME= | cut -c 6-
TARGET=server
APPNAME=server

export GOPATH="$CURDIR"

cd src/$APPNAME

echo 'get and build'
go get -v

go build -o appinhouse

echo 'package...'

case "$ACTION" in

    install)
        bee pack -o "$HOME"  -exr pack.sh -exr server -exr test.conf 
        ;;
    update)
	    bee pack -o "$HOME"  -exr bin -exr server -exr conf 
        ;;
    *)
        usage
        ;;
esac

echo 'finished'
